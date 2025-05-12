package kafkaconsumer

import (
	"context"

	"github.com/optique-dev/core"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	Ignite() error
	Stop() error
	// Add more methods here
	HandleMessage(message *kafka.Message) error
}

type kafkaConsumer struct {
	config *Config
	r      *kafka.Reader
}

func NewKafkaConsumer(config *Config) (*kafkaConsumer, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   config.Brokers,
		Topic:     config.Topic,
		Partition: config.Partition,
		MaxBytes:  10e6,
		GroupID:   config.GroupID,
	})
	return &kafkaConsumer{config: config,
		r: r}, nil
}

// Ignite starts the kafka consumer and listens for messages
// messages are handled Asynchronously
func (m *kafkaConsumer) Ignite() error {
	for {
		message, err := m.r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		go func(message *kafka.Message) {
			err := m.HandleMessage(message)
			core.Error(err.Error())
		}(&message)
	}
	return nil
}

func (m *kafkaConsumer) Stop() error {
	if err := m.r.Close(); err != nil {
		return err
	}
	return nil
}

func (m *kafkaConsumer) HandleMessage(message *kafka.Message) error {
	panic("implement me")
}
