package repository

import (
	"cek/model/domain"

	"gorm.io/gorm"
)

type TestRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Tests
}
