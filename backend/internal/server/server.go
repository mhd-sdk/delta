package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"sync"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/delta/internal/client"
	"github.com/delta/internal/pb"
	"github.com/delta/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket"
	"github.com/kr/pretty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	RoutingClient    *alpaca.Client
	MarketDataClient *marketdata.Client
	fiberServer      *fiber.App
	clients          []*client.Client
	mu               sync.Mutex
	slaves           map[string]pb.StateServiceClient
	activeStreams    map[string]context.CancelFunc // Track active log streams
}

func New() *Server {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	marketDataUrl := os.Getenv("MARKET_DATA_URL")
	routingUrl := os.Getenv("ROUTING_URL")

	RoutingClient := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   routingUrl,
	})

	MarketDataClient := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   marketDataUrl,
	})

	fiberServer := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	fiberServer.Use(logger.New(logger.Config{}))
	fiberServer.Use(cors.New())

	s := &Server{
		clients:          []*client.Client{},
		MarketDataClient: MarketDataClient,
		RoutingClient:    RoutingClient,
		fiberServer:      fiberServer,
		slaves:           make(map[string]pb.StateServiceClient),
		activeStreams:    make(map[string]context.CancelFunc),
	}
	initHandlers(s)
	go s.discoverSlaves()
	return s
}

func (s *Server) discoverSlaves() {
	addr, _ := net.ResolveUDPAddr("udp", ":9999")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, remoteAddr, _ := conn.ReadFromUDP(buf)
		slaveID := string(buf[:n])
		s.mu.Lock()
		if _, ok := s.slaves[slaveID]; !ok {
			connStr := fmt.Sprintf("%s:50051", remoteAddr.IP.String())
			clientConn, err := grpc.Dial(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Println("Impossible de se connecter à l'esclave:", err)
			} else {
				s.slaves[slaveID] = pb.NewStateServiceClient(clientConn)
				log.Println("Esclave découvert et ajouté:", slaveID, connStr)
			}
		}
		s.mu.Unlock()
	}
}

func (s *Server) SendCommand(slaveID, cmd string) {
	s.mu.Lock()
	client, ok := s.slaves[slaveID]
	s.mu.Unlock()
	if !ok {
		log.Println("Esclave non trouvé:", slaveID)
		return
	}
	_, err := client.Control(context.Background(), &pb.ControlRequest{Command: cmd})
	if err != nil {
		log.Println("Erreur en envoyant la commande:", err)
	} else {
		log.Println("Commande envoyée à", slaveID, ":", cmd)
	}
}

func (s *Server) ServeAPI() error {
	slog.Info("Serving api on localhost:3000")
	return s.fiberServer.Listen(":3000")
}

func (s *Server) TestCredentials() error {
	_, err := s.RoutingClient.GetAccount()
	return err
}

func (s *Server) AddClient(c *client.Client) {
	s.clients = append(s.clients, c)
}

func (s *Server) RemoveClient(c *client.Client) {
	for i, client := range s.clients {
		if client == c {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}
}

func (s *Server) GetAssets() (tickers []alpaca.Asset, err error) {
	if s.RoutingClient == nil {
		return nil, errors.New("not logged in")
	}

	assets, err := s.RoutingClient.GetAssets(alpaca.GetAssetsRequest{
		Status:     "active",
		AssetClass: "us_equity",
	})
	if err != nil {
		slog.Error("error fetching assets", "error", err.Error())
		return nil, err
	}
	return assets, nil
}

func (s *Server) GetAccount() (*alpaca.Account, error) {
	if s.RoutingClient == nil {
		slog.Error("not logged in")
		return nil, errors.New("not logged in")
	}
	acct, err := s.RoutingClient.GetAccount()
	if err != nil {
		return nil, err
	}
	return acct, nil
}

type GetCandlesticksConfig struct {
	Ticker    string          `json:"symbol"`
	Start     time.Time       `json:"start"`
	End       time.Time       `json:"end"`
	Timeframe model.Timeframe `json:"timeframe"`
}

type data struct {
	Data GetCandlesticksConfig `json:"data"`
}

func (s *Server) UpdateState() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.UpdateStateRequest
		if err := c.BodyParser(&req); err != nil {
			slog.Error("error parsing request body", "error", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		s.SendCommand(req.ID, req.Status)
		return c.JSON(fiber.Map{"status": "command sent"})
	}
}

func (s *Server) GetCandlesticks() fiber.Handler {
	return func(c *fiber.Ctx) error {

		querydata := data{}
		if err := c.BodyParser(&querydata); err != nil {
			slog.Error("error parsing query parameters", "error", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query parameters"})
		}
		config := querydata.Data

		// return c.JSON(chat)
		if s.MarketDataClient == nil {
			slog.Error("not logged in to market data client")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "not logged in"})
		}
		pretty.Println(config)
		data, err := s.MarketDataClient.GetBars(config.Ticker, marketdata.GetBarsRequest{
			TimeFrame: marketdata.TimeFrame{
				N:    config.Timeframe.N,
				Unit: marketdata.TimeFrameUnit(config.Timeframe.Unit),
			},
			Sort:  marketdata.SortAsc,
			Start: config.Start,
			// End:   config.End,
			Feed: marketdata.IEX,
		})
		if err != nil {
			slog.Error("error fetching assets", "error", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch candlesticks"})
		}

		return c.JSON(data)
	}

}

func (s *Server) subscribeToAllSlaveLogs(ctx context.Context, logChan chan<- *pb.LogEvent) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Initial subscription
	s.checkAndSubscribeToNewSlaves(ctx, logChan)

	for {
		select {
		case <-ctx.Done():
			// Cancel all active streams
			s.mu.Lock()
			for _, cancel := range s.activeStreams {
				cancel()
			}
			s.activeStreams = make(map[string]context.CancelFunc)
			s.mu.Unlock()
			return
		case <-ticker.C:
			s.checkAndSubscribeToNewSlaves(ctx, logChan)
		}
	}
}

func (s *Server) checkAndSubscribeToNewSlaves(ctx context.Context, logChan chan<- *pb.LogEvent) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check for new slaves that don't have active streams
	for slaveID, client := range s.slaves {
		if _, hasStream := s.activeStreams[slaveID]; !hasStream {
			// Create a new context for this slave's stream
			streamCtx, cancel := context.WithCancel(ctx)
			s.activeStreams[slaveID] = cancel

			go func(id string, c pb.StateServiceClient, sCtx context.Context) {
				s.subscribeToSlaveLogs(sCtx, id, c, logChan)
				// Clean up when stream ends
				s.mu.Lock()
				delete(s.activeStreams, id)
				s.mu.Unlock()
			}(slaveID, client, streamCtx)

			log.Printf("Started log subscription for slave: %s", slaveID)
		}
	}
}

func (s *Server) subscribeToSlaveLogs(ctx context.Context, slaveID string, client pb.StateServiceClient, logChan chan<- *pb.LogEvent) {
	stream, err := client.SubscribeLogs(ctx, &pb.LogSubscriptionRequest{
		SlaveId: slaveID,
	})
	if err != nil {
		log.Printf("Erreur lors de l'abonnement aux logs de %s: %v", slaveID, err)
		return
	}

	log.Printf("Successfully subscribed to logs for slave: %s", slaveID)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Log subscription cancelled for slave: %s", slaveID)
			return
		default:
			logEvent, err := stream.Recv()
			if err != nil {
				log.Printf("Erreur lors de la réception de logs de %s: %v", slaveID, err)
				return
			}
			fmt.Printf("Received log from %s: %s\n", slaveID, logEvent.Message)

			// broadcast the log event to all clients

			for _, client := range s.clients {
				if client.Conn == nil {
					continue
				}
				client.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[%s] %s: %s", slaveID, logEvent.Level, logEvent.Message)))
			}

			select {
			case logChan <- logEvent:
			case <-ctx.Done():
				return
			default:
				// Channel is full, skip this log event
				log.Printf("Log channel full, skipping event from %s", slaveID)
			}
		}
	}
}
