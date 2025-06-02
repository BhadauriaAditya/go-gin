package email

import (
	"context"
	"encoding/json"
	"log"

	"go-gin/internal/domain/models"
	"go-gin/internal/infrastructure/kafka"
)

type EmailConsumer struct {
	consumer kafka.Consumer
}

type UserCreatedEvent struct {
	User models.User `json:"user"`
}

func NewEmailConsumer(brokers []string) *EmailConsumer {
	consumer := kafka.NewConsumer(brokers, kafka.TopicUserCreated, "email-consumer-group")
	return &EmailConsumer{
		consumer: consumer,
	}
}

func (c *EmailConsumer) Start(ctx context.Context) {
	for {
		msg, err := c.consumer.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var event UserCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		if err := c.sendWelcomeEmail(event.User); err != nil {
			log.Printf("Error sending welcome email: %v", err)
			continue
		}
	}
}

func (c *EmailConsumer) sendWelcomeEmail(user models.User) error {
	// TODO: Implement email sending logic
	log.Printf("Sending welcome email to user: %s", user.Email)
	return nil
}

func (c *EmailConsumer) Close() error {
	return c.consumer.Close()
}
