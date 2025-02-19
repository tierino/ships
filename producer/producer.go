package producer

import (
	"encoding/json"
	"fmt"
)

type Producer struct{}

func New() *Producer {
	// init
	return &Producer{}
}

func (s *Producer) Send(topic string, msg interface{}) error {
	out, _ := json.Marshal(msg)
	fmt.Printf("%s: %s\n", topic, out)
	return nil
}
