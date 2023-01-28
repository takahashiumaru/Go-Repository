package repository

import (
	"cek/model/domain"

	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(db *gorm.DB, session *domain.Session) *domain.Session
	FindByRefreshUUID(db *gorm.DB, refreshUUID *string) domain.Session
	DeleteByRefreshUUID(db *gorm.DB, refreshUUID *string)
}
