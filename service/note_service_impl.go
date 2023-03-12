package service

import (
	"cek/auth"
	"cek/helper"
	"cek/model/domain"
	"cek/model/web"
	"cek/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewNoteService(
	note repository.NoteRepository,
	db *gorm.DB,
	validate *validator.Validate,
) NoteService {
	return &NoteServiceImpl{
		NoteRepository: note,
		DB:             db,
		Validate:       validate,
	}
}

func (service *NoteServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string) []web.NoteResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	notes := service.NoteRepository.FindAll(tx, filters)
	return notes.ToNoteResponses()
}

func (service *NoteServiceImpl) Create(auth *auth.AccessDetails, request *web.NoteCreateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	note := &domain.Note{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		Subject: request.Subject,
		Note:    request.Note,
	}
	notes := service.NoteRepository.Create(tx, note)
	return notes.ToNoteResponse()
}

func (service *NoteServiceImpl) Update(auth *auth.AccessDetails, noteID *int, request *web.NoteUpdateRequest) web.NoteResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	note := &domain.Note{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		ID:      uint(*noteID),
		Subject: request.Subject,
		Note:    request.Note,
	}
	note = service.NoteRepository.Update(tx, note)
	return note.ToNoteResponse()
}

func (service *NoteServiceImpl) Delete(auth *auth.AccessDetails, noteID *int) {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	service.NoteRepository.Delete(tx, noteID, &auth.UserID)
	defer helper.CommitOrRollback(tx)
}
