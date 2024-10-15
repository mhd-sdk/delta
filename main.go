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
	"github.com/lmittmann/tint"
	"google.golang.org/protobuf/proto"
)

var (
	addr = flag.String("addr", " wss://rituz00100.rithmic.com:443", "address of the server")
)

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
	if !slices.Contains(response.GetSystemName(), "Rithmic Test") {
		slog.Error("Rithmic test server unreachable, exiting...")
		os.Exit(1)
	}

	usr := "mhdi.seddik@gmail.com"
	pwd := "lDIKLQCX"

	slog.Info("Rithmic test server found, logging in as " + usr + "...")

	loginRequest := rti.RequestLogin{
		User:     proto.String(usr),
		Password: proto.String(pwd),
	}

	data, err := proto.Marshal(loginRequest.ProtoReflect().Interface())
	if err != nil {
		log.Fatal("Error marshalling message:", err)
		return
	}

	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

}
