package repository

import (
	"cek/model/domain"

	"gorm.io/gorm"
)

type NoteRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Notes
	Create(db *gorm.DB, note *domain.Note) *domain.Note
	Update(db *gorm.DB, note *domain.Note) *domain.Note
	Delete(db *gorm.DB, noteID *int, deletedByID *uint)
}
