package repository

import (
	"errors"

	"cek/exception"
	"cek/helper"
	"cek/model/domain"

	"gorm.io/gorm"
)

type SessionRepositoryImpl struct {
}

func NewSessionRepository() SessionRepository {
	return &SessionRepositoryImpl{}
}

func (repository *SessionRepositoryImpl) Create(db *gorm.DB, session *domain.Session) *domain.Session {
	err := db.Create(&session).Error
	helper.PanicIfError(err)
	return session
}

func (repository *SessionRepositoryImpl) FindByRefreshUUID(db *gorm.DB, refreshUUID *string) domain.Session {
	var session domain.Session
	result := db.Where(&domain.Session{
		RefreshUUID: *refreshUUID,
	}).First(&session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helper.PanicIfError(exception.ErrUnauthorized)
	}
	helper.PanicIfError(result.Error)
	return session
}

func (repository *SessionRepositoryImpl) DeleteByRefreshUUID(db *gorm.DB, refreshUUID *string) {
	err := db.Unscoped().Where(&domain.Session{
		RefreshUUID: *refreshUUID,
	}).Delete(&domain.Session{}).Error
	helper.PanicIfError(err)
}
