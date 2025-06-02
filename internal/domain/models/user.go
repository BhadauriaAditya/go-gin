package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	AvatarURL string `json:"avatar_url,omitempty"`
	MoodTag   string `json:"mood_tag,omitempty"`
}
