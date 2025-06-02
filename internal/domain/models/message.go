package models

import (
	"gorm.io/gorm"
	"time"
)

// *User when the sender is not specified, it can be nil enables lazy and eger laoding for the model.
//*time vs time -- use time when you need to represent a specific point in time, use *time when you want to represent an optional time value.

type Message struct {
	gorm.Model
	SenderID     *uint      `gorm:"index" json:"sender_id,omitempty"`
	Sender       *User      `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	RecipientID  uint       `gorm:"not null;index" json:"recipient_id"`
	Recipient    *User      `gorm:"foreignKey:RecipientID" json:"recipient,omitempty"`
	Content      string     `gorm:"type:text;not null" json:"content"`
	IsAnonymous  bool       `gorm:"default:false" json:"is_anonymous"`
	IsPublic     bool       `gorm:"default:false" json:"is_public"`
	CanSeeOnce   bool       `gorm:"default:false" json:"can_see_once"`
	AutoExpireAt *time.Time `json:"auto_expire_at,omitempty"`
	VisibleUntil *time.Time `json:"visible_until,omitempty"`
}
