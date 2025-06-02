package models

import (
	"gorm.io/gorm"
	"time"
)

type MessageVisibilityLog struct {
	gorm.Model
	MessageID *uint      `gorm:"type:uuid;not null;index" json:"message_id"`
	Message   *Message   `gorm:"foreignKey:MessageID" json:"-"`
	UserID    *uint      `gorm:"type:uuid;not null;index" json:"user_id"`
	User      *User      `gorm:"foreignKey:UserID" json:"-"`
	ViewedAt  *time.Time `gorm:"autoCreateTime" json:"viewed_at"`
}
