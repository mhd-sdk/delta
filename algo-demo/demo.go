package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"slave/pb"
	"time"

	"google.golang.org/grpc"
)

type SlaveServer struct {
	pb.UnimplementedStateServiceServer
	id        string
	state     string // running or stopped
	logStream chan *pb.LogEvent
}

func NewSlaveServer(id string) *SlaveServer {
	return &SlaveServer{
		id:        id,
		state:     "stopped",
		logStream: make(chan *pb.LogEvent, 100),
	}

}

func (s *SlaveServer) Control(ctx context.Context, req *pb.ControlRequest) (*pb.ControlResponse, error) {
	s.state = req.Command
	s.sendLogEvent("info", fmt.Sprintf("Commande reçue: %s", req.Command))

	if s.state == "start" {
		s.sendLogEvent("info", fmt.Sprintf("Esclave %s démarré", s.id))
		s.state = "running"
	} else if s.state == "stop" {
		s.sendLogEvent("info", fmt.Sprintf("Esclave %s arrêté", s.id))
		s.state = "stopped"
	} else {
		s.sendLogEvent("error", fmt.Sprintf("Commande inconnue: %s", s.state))
		return nil, fmt.Errorf("commande inconnue: %s", s.state)
	}
	return &pb.ControlResponse{Ack: true}, nil
}

func (s *SlaveServer) SubscribeLogs(req *pb.LogSubscriptionRequest, stream pb.StateService_SubscribeLogsServer) error {
	log.Printf("Client abonné aux logs pour esclave %s", s.id)
	// Send initial log event
	s.sendLogEvent("info", s.state)

	for {
		select {
		case logEvent := <-s.logStream:
			if err := stream.Send(logEvent); err != nil {
				log.Printf("Erreur envoi log stream: %v", err)
				return err
			}
		case <-stream.Context().Done():
			log.Printf("Client déconnecté du stream de logs")
			return nil
		}
	}
}

func (s *SlaveServer) sendLogEvent(level, message string) {
	logEvent := &pb.LogEvent{
		SlaveId:   s.id,
		Message:   message,
		Timestamp: time.Now().Unix(),
		Level:     level,
	}

	select {
	case s.logStream <- logEvent:
	default:
		// Channel is full, skip this log event
		log.Printf("Log channel full, skipping log event: %s", message)
	}

	log.Printf("[%s] %s: %s", level, s.id, message)
}

func (s *SlaveServer) runAlgorithm() {

	for {
		fmt.Println(s.state)
		if s.state == "running" {
			s.sendLogEvent("debug", "running")
			type BuyEvent struct {
				Asset    string  `json:"asset"`
				Price    float64 `json:"price"`
				Quantity int     `json:"quantity"`
			}
			randombuyEvent := BuyEvent{
				Asset:    "AAPL",
				Price:    math.Round(150.0*100) / 100, // Simulate a price
				Quantity: 1,                           // Simulate a quantity
			}
			// struct to json string marshal
			data, err := json.Marshal(randombuyEvent)
			if err != nil {
				panic(err)
			}

			s.sendLogEvent("info", string(data))
		}
		if s.state == "stopped" {
			s.sendLogEvent("info", "stopped")
		}
		time.Sleep(5 * time.Second) // Simulate algorithm work
	}

}

func (s *SlaveServer) announcePresence() {
	addr, _ := net.ResolveUDPAddr("udp", "255.255.255.255:9999")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	for {
		conn.Write([]byte(s.id))
		time.Sleep(5 * time.Second)
	}
}

func main() {
	id := "Demo"
	slave := NewSlaveServer(id)

	go slave.announcePresence()

	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
	pb.RegisterStateServiceServer(grpcServer, slave)

	// Send startup log only once
	log.Println("Slave gRPC server started on :50051")

	go slave.runAlgorithm()
	grpcServer.Serve(lis)
}
