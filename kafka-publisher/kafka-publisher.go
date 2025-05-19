package kafkapublisher

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher interface {
	Setup() error
	Shutdown() error
	Write(message []byte) error
	// Add more methods here
}

type kafkaPublisher struct {
	conn *kafka.Writer
}

func NewKafkaPublisher(kafka_config Config) (*kafkaPublisher, error) {
	writer := kafka.Writer{
		Addr:     kafka.TCP(kafka_config.Brokers...),
		Topic:    kafka_config.Topic,
		Balancer: &kafka.RoundRobin{},
	}
	return &kafkaPublisher{conn: &writer}, nil
}

func (m *kafkaPublisher) Setup() error {
	return nil
}

func (m *kafkaPublisher) Write(message []byte) error {
	err := m.conn.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *kafkaPublisher) Shutdown() error {
	return m.conn.Close()
}
