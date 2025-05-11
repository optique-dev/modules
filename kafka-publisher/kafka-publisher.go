package kafkapublisher

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	Setup() error
	Shutdown() error
	Write(message []byte) error
	// Add more methods here
}

type kafkaConsumer struct {
	conn *kafka.Conn
}

func NewKafkaConsumer(kafka_config Config) (*kafkaConsumer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafka_config.Brokers[0], kafka_config.Topic, 0)
	if err != nil {
		return nil, err
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return &kafkaConsumer{conn: conn}, nil
}

func (m *kafkaConsumer) Setup() error {
	return nil
}

func (m *kafkaConsumer) Write(message []byte) error {
	m.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := m.conn.WriteMessages(kafka.Message{
		Value: message,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *kafkaConsumer) Shutdown() error {
	return m.conn.Close()
}
