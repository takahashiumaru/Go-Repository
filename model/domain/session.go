package domain

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	// Required Fields
	gorm.Model

	// Fields
	RefreshUUID   string    `gorm:"not null;unique"`
	UserID        uint      `gorm:"not null"`
	UserAgent     string    `gorm:"not null"`
	RemoteAddress string    `gorm:"not null"`
	Expired       time.Time `gorm:"not null"`
}
