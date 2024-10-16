package services

import (
	"context"
	"delta/pkg/generated/rti"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type BaseService struct {
	wsClient *websocket.Conn
}

type Service interface {
	Connect(ctx context.Context) error
	Close() error
}

type ConnectionArgs struct {
	Url        string
	SystemName string
	User       string
	Password   string
}

func (s *BaseService) Connect(connectionArgs ConnectionArgs, infraType rti.RequestLogin_SysInfraType) error {
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
		return err
	}

	conn, _, err := websocket.DefaultDialer.Dial(connectionArgs.Url, nil)
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
	return nil
}
