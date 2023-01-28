package repository

import (
	"errors"

	"cek/exception"
	"cek/helper"
	"cek/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(db *gorm.DB, user *domain.User) *domain.User {
	err := db.Create(&user).Error
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) FindByID(db *gorm.DB, id *int) domain.User {
	var user domain.User
	result := db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helper.PanicIfError(exception.ErrUnauthorized)
	}
	helper.PanicIfError(result.Error)
	return user
}

func (repository *UserRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Users {
	var users domain.Users
	tx := db.Model(&domain.User{})
	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	result := tx.Find(&users)
	helper.PanicIfError(result.Error)
	return users
}

func (repository *UserRepositoryImpl) FindByNip(db *gorm.DB, nip *string) domain.User {
	var user domain.User
	result := db.Where("nip = ? AND deleted_at IS NULL", *nip).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helper.PanicIfError(exception.ErrUnauthorized)
	}
	helper.PanicIfError(result.Error)
	return user
}

func (repository *UserRepositoryImpl) Delete(db *gorm.DB, id *int) {
	err := db.Delete(&domain.User{}, id).Error
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) Update(db *gorm.DB, user *domain.User) *domain.User {
	err := db.Updates(&user).Error
	helper.PanicIfError(err)
	return user
}
