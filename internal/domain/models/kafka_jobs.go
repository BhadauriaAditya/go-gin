package models

import (
	"gorm.io/gorm"
)

type kafkaJobs struct {
	gorm.Model
	MessageID *uint    `gorm:"type:uuid;not null;index" json:"message_id"`
	Message   *Message `gorm:"foreignKey:MessageID" json:"-"`
	EventType string   `gorm:"not null" json:"event_type"` // e.g., "message_sent", "message_viewed", "message_deleted"
	Status    string   `gorm:"not null" json:"status"`     // e.g., "pending", "processing", "completed", "failed"
}
