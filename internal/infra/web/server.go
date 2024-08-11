package web

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/sandronister/go_broker/pkg/payload"
	"github.com/sandronister/go_broker/pkg/ports"
	"github.com/sandronister/socket-go/config"
)

type Server struct {
	name   string
	host   string
	port   string
	topic  string
	conn   net.Conn
	broker ports.IBroker
}

func NewServer(device *config.Device, broker ports.IBroker) *Server {
	return &Server{
		name:   device.Name,
		host:   device.Host,
		port:   device.Port,
		topic:  device.Topic,
		broker: broker,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.host+":"+s.port)

	fmt.Println("Server started at", s.host+":"+s.port, "for", s.name)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		s.conn, err = listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go s.HandleConnection()

	}
}

func (s *Server) GetPayload(msg string) *payload.Message {
	return &payload.Message{
		TopicPartition: s.topic,
		Value:          []byte(msg),
		Key:            []byte(uuid.New().String()),
		Timestamp:      time.Unix(time.Now().Unix(), 0),
	}
}

func (s *Server) HandleConnection() {
	defer s.conn.Close()
	reader := bufio.NewReader(s.conn)

	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
			return
		}

		if msg != "" {
			payload := s.GetPayload(msg)
			s.broker.Produce(payload)
		}
	}

}
