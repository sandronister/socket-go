package uruptela

import (
	"fmt"
	"log"
	"time"

	"gitlab.com/gobrax-dev/gobrax-tool/broker/types"
)

func (u *UseRuptela) SendMessage(msg []byte) {
	payload := &types.Message{
		Value:     msg,
		Key:       []byte(fmt.Sprintf("%d", time.Now().Unix())),
		Timestamp: time.Now(),
	}

	if len(msg) != 0 {
		err := u.broker.Producer(payload, 1)

		if err != nil {
			log.Println(err)
		}
	}
}
