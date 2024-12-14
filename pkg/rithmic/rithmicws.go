package rithmic

import (
	"delta/pkg/generated/rti"
	"log"
	"log/slog"
	"os"
	"slices"

	"github.com/gorilla/websocket"
	"github.com/kr/pretty"
	"google.golang.org/protobuf/proto"
)

type RithmicWS struct {
	wsClients (map[rti.RequestLogin_SysInfraType]*websocket.Conn)
}

type ConnectionArgs struct {
	Url        string
	User       string
	Password   string
	SystemName string
}

func New(connectionArgs ConnectionArgs) *RithmicWS {
	conn, _, err := websocket.DefaultDialer.Dial(connectionArgs.Url, nil)
	if err != nil {
		slog.Error("Error connecting to WebSocket", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	systemInfoRequest := rti.RequestRithmicSystemInfo{
		TemplateId: proto.Int32(16),
	}

	data, err := proto.Marshal(systemInfoRequest.ProtoReflect().Interface())
	if err != nil {
		slog.Error("Error marshalling", "message", err)
		os.Exit(1)
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		slog.Error("Error sending", "message", err)
		os.Exit(1)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		slog.Error("Error reading", "message", err)
		os.Exit(1)
	}

	var response rti.ResponseRithmicSystemInfo
	err = proto.Unmarshal(msg, &response)
	if err != nil {
		slog.Error("Error unmarshalling", "message", err)
		os.Exit(1)
	}

	pretty.Println(response.SystemName)

	if connectionArgs.SystemName == "" {
		connectionArgs.SystemName = DEFAULT_RITHMIC_SYSTEM_NAME
	}

	if !slices.Contains(response.GetSystemName(), connectionArgs.SystemName) {
		slog.Error("Error connecting to rithmic services, requested system not found", "system name", connectionArgs.SystemName)
		slog.Info("Here's a list of existing systems", "system names", response.GetSystemName())
		os.Exit(1)
	}

	conn.Close()

	client := &RithmicWS{
		wsClients: make(map[rti.RequestLogin_SysInfraType]*websocket.Conn),
	}

	for _, infraType := range AVAILABLE_RITHMIC_INFRA_TYPES {
		loginRequest := rti.RequestLogin{
			InfraType:       &infraType,
			TemplateId:      proto.Int32(10),
			TemplateVersion: proto.String("3.9"),
			AppName:         proto.String("DeltΔ"),
			AppVersion:      proto.String("1.0.0"),
			User:            proto.String(connectionArgs.User),
			Password:        proto.String(connectionArgs.Password),
			SystemName:      proto.String(connectionArgs.SystemName),
		}

		data, err := proto.Marshal(loginRequest.ProtoReflect().Interface())
		if err != nil {
			slog.Error("Error marshalling", "message", err)
			os.Exit(1)
		}

		conn, _, err := websocket.DefaultDialer.Dial(connectionArgs.Url, nil)
		if err != nil {
			slog.Error("Error connecting to WebSocket", "error", err)
			os.Exit(1)
		}

		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			slog.Error("Error sending", "message", err)
			os.Exit(1)
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error reading", "message", err)
			os.Exit(1)
		}

		var loginResponse rti.ResponseLogin

		err = proto.Unmarshal(msg, &loginResponse)
		if err != nil {
			slog.Error("Error unmarshalling", "message", err)
			os.Exit(1)
		}

		client.wsClients[infraType] = conn
	}
	return client
}

func (r *RithmicWS) Close() {
	for _, conn := range r.wsClients {
		conn.Close()
	}
}

type MarketDataLastTradeHandler func(rti.LastTrade) error

func (r *RithmicWS) SubscribeMarketDataLastTrade(symbol string, exchange string, handler MarketDataLastTradeHandler) {
	go func() {
		lt := uint32(rti.RequestMarketDataUpdate_LAST_TRADE)
		conn := r.wsClients[rti.RequestLogin_TICKER_PLANT]

		rq := rti.RequestMarketDataUpdate{
			TemplateId: proto.Int32(100),
			Symbol:     proto.String(symbol),
			Exchange:   proto.String(exchange),
			Request:    rti.RequestMarketDataUpdate_SUBSCRIBE.Enum(),
			UpdateBits: &lt,
		}

		data, err := proto.Marshal(rq.ProtoReflect().Interface())
		if err != nil {
			slog.Error("Error marshalling", "message", err)
			os.Exit(1)
		}

		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		var marketDataResponse rti.ResponseMarketDataUpdate

		err = proto.Unmarshal(msg, &marketDataResponse)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			return
		}

		for {
			_, msg, err = conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			var lastTrade rti.LastTrade

			err = proto.Unmarshal(msg, &lastTrade)
			if err != nil {
				log.Println("Error unmarshalling message:", err)
				return
			}

			handler(lastTrade)
		}
	}()
}
