package rithmic

import (
	"delta/pkg/generated/rti"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type RithmicWS struct {
	wsClients       (map[rti.RequestLogin_SysInfraType]*websocket.Conn)
	heartbeatClient *websocket.Conn
	connectionArgs  ConnectionArgs
}

type ConnectionArgs struct {
	Url        string
	User       string
	Password   string
	SystemName string
}

func (r *RithmicWS) Login() error {
	for _, infraType := range AVAILABLE_RITHMIC_INFRA_TYPES {
		loginRequest := rti.RequestLogin{
			InfraType:       &infraType,
			TemplateId:      proto.Int32(10),
			TemplateVersion: proto.String("3.9"),
			AppName:         proto.String("mese:Delta"),
			AppVersion:      proto.String("1.0.0"),
			User:            proto.String(r.connectionArgs.User),
			Password:        proto.String(r.connectionArgs.Password),
			SystemName:      proto.String(r.connectionArgs.SystemName),
		}

		data, err := proto.Marshal(loginRequest.ProtoReflect().Interface())
		if err != nil {
			return err
		}

		conn, _, err := websocket.DefaultDialer.Dial(r.connectionArgs.Url, nil)
		if err != nil {
			return err
		}

		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			return err
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		var loginResponse rti.ResponseLogin

		err = proto.Unmarshal(msg, &loginResponse)
		if err != nil {
			return err
		}

		fmt.Println(loginResponse.String())

		r.wsClients[infraType] = conn
	}

	return nil
}

func New(connectionArgs ConnectionArgs) (*RithmicWS, error) {
	client := &RithmicWS{
		wsClients:      make(map[rti.RequestLogin_SysInfraType]*websocket.Conn),
		connectionArgs: connectionArgs,
	}

	return client, nil
}

func (r *RithmicWS) ListSystems() (*rti.ResponseRithmicSystemInfo, error) {
	conn, _, err := websocket.DefaultDialer.Dial(r.connectionArgs.Url, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	systemInfoRequest := rti.RequestRithmicSystemInfo{
		TemplateId: proto.Int32(16),
	}

	data, err := proto.Marshal(systemInfoRequest.ProtoReflect().Interface())
	if err != nil {
		return nil, err
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return nil, err
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	var response rti.ResponseRithmicSystemInfo

	err = proto.Unmarshal(msg, &response)
	if err != nil {
		return nil, err
	}

	conn.Close()

	return &response, nil
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

func (r *RithmicWS) ListProducts() ([]*rti.ResponseProductCodes, error) {
	conn := r.wsClients[rti.RequestLogin_TICKER_PLANT]
	if conn == nil {
		return nil, fmt.Errorf("no connection to rithmic")
	}

	rq := rti.RequestProductCodes{
		TemplateId: proto.Int32(111),
		Exchange:   proto.String("CME"),
	}

	data, err := proto.Marshal(rq.ProtoReflect().Interface())
	if err != nil {
		slog.Error("Error marshalling", "message", err)
		return nil, err
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		slog.Error("Error sending message", "error", err)
		return nil, err
	}

	var productCodes []*rti.ResponseProductCodes
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return nil, err
		}

		var productCodesResponse rti.ResponseProductCodes

		err = proto.Unmarshal(msg, &productCodesResponse)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			return nil, err
		}

		// Check for the presence of fields to determine if more messages are expected
		sequenceFinished := len(productCodesResponse.RqHandlerRpCode) == 0 && len(productCodesResponse.RpCode) != 0
		if sequenceFinished {
			break // No more messages to receive
		}

		productCodes = append(productCodes, &productCodesResponse)

	}

	return productCodes, nil
}

func (r *RithmicWS) SearchSymbols(productCode string) ([]*rti.ResponseSearchSymbols, error) {
	conn := r.wsClients[rti.RequestLogin_TICKER_PLANT]

	rq := rti.RequestSearchSymbols{
		TemplateId: proto.Int32(109),
		SearchText: proto.String("MES"),
		// Pattern:    rti.RequestSearchSymbols_CONTAINS.Enum(),
		Pattern: rti.RequestSearchSymbols_EQUALS.Enum().Enum(),
	}

	data, err := proto.Marshal(rq.ProtoReflect().Interface())
	if err != nil {
		slog.Error("Error marshalling", "message", err)
		return nil, err
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Println("Error sending message:", err)
		return nil, err
	}

	var symbols []*rti.ResponseSearchSymbols
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return nil, err
		}

		var response rti.ResponseSearchSymbols

		err = proto.Unmarshal(msg, &response)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			return nil, err
		}

		// Check for the presence of fields to determine if more messages are expected
		sequenceFinished := len(response.RqHandlerRpCode) == 0 && len(response.RpCode) != 0
		if sequenceFinished {
			break // No more messages to receive
		}

		symbols = append(symbols, &response)

	}

	return symbols, nil
}

func (r *RithmicWS) GetInstrument(symbolName string) ([]*rti.ResponseGetInstrumentByUnderlying, error) {
	conn := r.wsClients[rti.RequestLogin_TICKER_PLANT]

	rq := rti.RequestGetInstrumentByUnderlying{
		TemplateId:       proto.Int32(102),
		UnderlyingSymbol: proto.String(symbolName),
	}

	data, err := proto.Marshal(rq.ProtoReflect().Interface())
	if err != nil {
		slog.Error("Error marshalling", "message", err)
		return nil, err
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Println("Error sending message:", err)
		return nil, err
	}

	var instruments []*rti.ResponseGetInstrumentByUnderlying
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return nil, err
		}

		var instrumentResponse rti.ResponseGetInstrumentByUnderlying

		err = proto.Unmarshal(msg, &instrumentResponse)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			return nil, err
		}

		fmt.Println(instrumentResponse.String())

		instruments = append(instruments, &instrumentResponse)

		// Check for the presence of fields to determine if more messages are expected
		sequenceFinished := len(instrumentResponse.RqHandlerRpCode) == 0 && len(instrumentResponse.RpCode) != 0
		if sequenceFinished {
			break // No more messages to receive
		}

	}

	return instruments, nil
}
