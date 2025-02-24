package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Config struct {
	BootstrapServers string
}

type Producer struct {
	kp *kafka.Producer
}

func New(conf *Config) (*Producer, error) {
	kp, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.BootstrapServers,
	})
	if err != nil {
		fmt.Println("coconut")
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}
	return &Producer{kp}, nil
}

func (p *Producer) Send(topic string, msg []byte) error {
	err := p.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)
	return err
}

func (p *Producer) Disconnect() {
	p.kp.Close()
}
