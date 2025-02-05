package server

import (
	"log/slog"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/delta/internal/client"
	"github.com/delta/internal/filelogger"
	"github.com/delta/internal/worker/scanner"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	fileLogger       *filelogger.FileLogger
	routingClient    *alpaca.Client
	marketDataClient *marketdata.Client
	fiberServer      *fiber.App
	scanner          *scanner.Scanner
	clients          []*client.Client
}

func New() *Server {
	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	marketDataUrl := os.Getenv("MARKET_DATA_URL")
	routingUrl := os.Getenv("ROUTING_URL")

	fileLoger, err := filelogger.NewLogger("delta.log")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	routingClient := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   routingUrl,
	})

	marketDataClient := marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    apiKey,
		APISecret: secretKey,
		BaseURL:   marketDataUrl,
	})

	fiberServer := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	fiberServer.Use(logger.New(logger.Config{}))

	scanner := scanner.New(scanner.ScannerConfig{})

	s := &Server{
		clients:          []*client.Client{},
		marketDataClient: marketDataClient,
		routingClient:    routingClient,
		fiberServer:      fiberServer,
		scanner:          scanner,
		fileLogger:       fileLoger,
	}
	initHandlers(s)

	s.scanner.Subscribe(fileLoger)

	return s
}

func (s *Server) ServeAPI() error {
	slog.Info("Serving api on localhost:3000")
	return s.fiberServer.Listen(":3000")
}

func (s *Server) StartScanner() {
	slog.Info("Starting scanner worker")
	go s.scanner.Start()
}

func (s *Server) TestCredentials() error {
	_, err := s.routingClient.GetAccount()
	return err
}

func (s *Server) AddClient(c *client.Client) {
	s.clients = append(s.clients, c)
	s.scanner.Subscribe(c)
}

func (s *Server) RemoveClient(c *client.Client) {
	for i, client := range s.clients {
		if client == c {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}
	s.scanner.Unsubscribe(c)
}
