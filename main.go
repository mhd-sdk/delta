package main

import (
	"delta/pkg/generated/rti"
	"flag"
	"log"
	"log/slog"
	"os"
	"slices"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kr/pretty"
	"github.com/lmittmann/tint"
	"google.golang.org/protobuf/proto"
)

var (
	addr = flag.String("addr", " wss://rituz00100.rithmic.com:443", "address of the server")
)

const RITHMIC_SYSTEM_NAME = "Rithmic Test"

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
	slog.Info("Starting DeltΔ...")

	url := "wss://rituz00100.rithmic.com:443"

	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
		os.Exit(1)
	}
	defer conn.Close()

	slog.Info("Requesting system names...")

	systemInfoRequest := rti.RequestRithmicSystemInfo{
		TemplateId: proto.Int32(16),
	}

	data, err := proto.Marshal(systemInfoRequest.ProtoReflect().Interface())
	if err != nil {
		slog.Error("Error marshalling", "message", err)
		return
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		slog.Error("Error sending", "message", err)
		os.Exit(1)
	}

	// Listen for incoming messages

	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading", "message", err)
		os.Exit(1)
	}
	// decode ResponseRithmicSystemInfo
	var response rti.ResponseRithmicSystemInfo
	err = proto.Unmarshal(msg, &response)
	if err != nil {
		log.Fatal("Error unmarshalling message:", err)
		return
	}
	if !slices.Contains(response.GetSystemName(), RITHMIC_SYSTEM_NAME) {
		slog.Error(RITHMIC_SYSTEM_NAME + " server unreachable, exiting...")
		os.Exit(1)
	}

	conn.Close()

	usr := "mhdi.seddik@gmail.com"
	pwd := "lDIKLQCX"

	slog.Info(RITHMIC_SYSTEM_NAME + " server found, logging in...")

	conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
		os.Exit(1)
	}

	infraType := rti.RequestLogin_TICKER_PLANT
	loginRequest := rti.RequestLogin{
		TemplateId:      proto.Int32(10),
		User:            proto.String(usr),
		Password:        proto.String(pwd),
		SystemName:      proto.String(RITHMIC_SYSTEM_NAME),
		AppName:         proto.String("DeltΔ"),
		InfraType:       &infraType,
		AppVersion:      proto.String("1.0.0"),
		TemplateVersion: proto.String("3.9"),
	}

	data, err = proto.Marshal(loginRequest.ProtoReflect().Interface())
	if err != nil {
		log.Fatal("Error marshalling message:", err)
		return
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	_, msg, err = conn.ReadMessage()
	if err != nil {
		log.Println("Error reading message:", err)
		return
	}

	var loginResponse rti.ResponseLogin

	err = proto.Unmarshal(msg, &loginResponse)
	if err != nil {
		log.Println("Error unmarshalling message:", err)
		return
	}

	slog.Info("Login successful")

	lastTrade := uint32(rti.RequestMarketDataUpdate_LAST_TRADE)
	marketDataRequest := rti.RequestMarketDataUpdate{
		TemplateId: proto.Int32(100),
		Symbol:     proto.String("ESZ4"),
		Exchange:   proto.String("CME"),
		Request:    rti.RequestMarketDataUpdate_SUBSCRIBE.Enum(),
		UpdateBits: &lastTrade,
	}

	data, err = proto.Marshal(marketDataRequest.ProtoReflect().Interface())
	if err != nil {
		log.Fatal("Error marshalling message:", err)
		return
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	_, msg, err = conn.ReadMessage()
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

		pretty.Println(lastTrade.String())
	}
}
