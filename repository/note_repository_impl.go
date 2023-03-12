package repository

import (
	"cek/exception"
	"cek/helper"
	"cek/model/domain"
	"strings"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
}

func NewNoteRepository() NoteRepository {
	return &NoteRepositoryImpl{}
}

func (repository *NoteRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string) domain.Notes {
	notes := domain.Notes{}
	tx := db.Model(&domain.Note{}).
		Joins("CreatedBy").
		Joins("UpdatedBy")
	err := tx.Find(&notes).Error
	helper.PanicIfError(err)

	return notes
}

func (repository *NoteRepositoryImpl) Create(db *gorm.DB, note *domain.Note) *domain.Note {

	err := db.Create(&note).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = &exception.ErrorSendToResponse{Err: "record already exists"}
		}
	}
	helper.PanicIfError(err)
	return note
}

func (repository *NoteRepositoryImpl) Update(db *gorm.DB, note *domain.Note) *domain.Note {
	err := db.Updates(&note).Error
	helper.PanicIfError(err)

	err = db.First(&note).Error
	helper.PanicIfError(err)

	// err = helper.CreateHistory(db, eventDetail, helper.HistoryUpdate, eventDetail.UpdatedByID)
	// helper.PanicIfError(err)

	return note
}

func (repository *NoteRepositoryImpl) Delete(db *gorm.DB, id *int, deletedByID *uint) {
	eventDetail := &domain.Note{}
	tx := db.First(eventDetail, id).Updates(&domain.Note{
		// Model:       gorm.Model{ID: uint(*id)},
		DeletedByID: deletedByID,
	})

	// Creating a history of the deleted event detail.
	// err := helper.CreateHistory(db, eventDetail, helper.HistoryDelete, *deletedByID)
	// helper.PanicIfError(err)

	// Deleting the event detail from the database.
	err := tx.Unscoped().Delete(eventDetail, id).Error
	helper.PanicIfError(err)
}
