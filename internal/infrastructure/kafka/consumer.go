package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Consumer defines the interface for Kafka consumers
type Consumer interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close() error
}

// BaseConsumer implements the basic Kafka consumer functionality
type BaseConsumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string, groupID string) *BaseConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &BaseConsumer{
		reader: reader,
	}
}

func (c *BaseConsumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return kafka.Message{}, fmt.Errorf("failed to read message: %w", err)
	}
	return msg, nil
}

func (c *BaseConsumer) Close() error {
	return c.reader.Close()
}
