package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (h *TcpHandler) SendMessage(msg []byte) {
	payload := &types.Message{
		Value:     []byte(msg),
		Key:       []byte(fmt.Sprintf("%d", i)),
		Timestamp: time.Now(),
	}

	if len(msg) != 0 {
		err := h.broker.Producer(payload)
		if err != nil {
			log.Println(err)
		}
	}
}
