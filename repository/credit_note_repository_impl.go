package repository

import (
	"cek/helper"
	"cek/model/domain"

	"gorm.io/gorm"
)

type TestRepositoryImpl struct {
}

func NewTestRepository() TestRepository {
	return &TestRepositoryImpl{}
}

func (repository *TestRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Tests {
	customers := domain.Tests{}
	tx := db.Model(&domain.Test{})
	err := tx.Find(&customers).Error
	helper.PanicIfError(err)

	return customers
}
