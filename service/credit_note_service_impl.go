package service

import (
	"cek/auth"
	"cek/helper"
	"cek/model/web"
	"cek/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TestServiceImpl struct {
	TestRepository repository.TestRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewTestService(
	test repository.TestRepository,
	db *gorm.DB,
	validate *validator.Validate,
) TestService {
	return &TestServiceImpl{
		TestRepository: test,
		DB:             db,
		Validate:       validate,
	}
}

func (service *TestServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.TestResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	customers := service.TestRepository.FindAll(tx, filters)
	return customers.ToTestResponses()
}
