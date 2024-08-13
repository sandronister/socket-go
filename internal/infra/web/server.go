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
	host      string
	port      string
	topic     string
	conn      net.Conn
	broker    ports.IBroker
	timeFlush int
}

func NewServer(conf *config.Conf, broker ports.IBroker) *Server {
	return &Server{
		host:      conf.SOCKET_HOST,
		port:      conf.SOCKET_PORT,
		topic:     conf.DEVICE_TOPIC,
		broker:    broker,
		timeFlush: conf.TIME_FLUSH * 1000,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.host+":"+s.port)

	fmt.Println("Server started at", s.host+":"+s.port, "for", s.topic)

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

	reader := bufio.NewReader(s.conn)
	defer s.conn.Close()
	for {

		msg, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
			return
		}

		if msg != "" {
			payload := s.GetPayload(msg)
			s.broker.Produce(payload, s.timeFlush)
		}
	}

}
