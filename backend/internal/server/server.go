package server

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/delta/internal/client"
	"github.com/delta/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kr/pretty"
)

type Server struct {
	RoutingClient    *alpaca.Client
	MarketDataClient *marketdata.Client
	fiberServer      *fiber.App
	clients          []*client.Client
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
	}
	initHandlers(s)

	return s
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
