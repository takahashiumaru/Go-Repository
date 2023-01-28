package repository

import (
	"cek/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(tx *gorm.DB, filters *map[string]string) domain.Users
	FindByID(tx *gorm.DB, id *int) domain.User
	FindByNip(tx *gorm.DB, nip *string) domain.User
	Create(tx *gorm.DB, user *domain.User) *domain.User
	Delete(tx *gorm.DB, id *int)
	Update(tx *gorm.DB, user *domain.User) *domain.User
}
