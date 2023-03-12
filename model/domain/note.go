package domain

import (
	"time"

	"cek/model/web"

	"gorm.io/gorm"
)

type Notes []Note
type Note struct {
	// Required Fields
	gorm.Model
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:""`
	CreatedByID uint      `gorm:""`
	UpdatedAt   time.Time `gorm:""`
	UpdatedByID uint      `gorm:""`
	DeletedByID *uint     `gorm:""`
	CreatedBy   User
	UpdatedBy   User

	// Fields
	Subject string `gorm:"size:500;not null"`
	Note    string `gorm:"size:8000;not null"`
}

func (note *Note) ToNoteResponse() web.NoteResponse {

	return web.NoteResponse{
		// Required Fields
		ID:          note.ID,
		CreatedByID: note.CreatedByID,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
		CreatedBy:   note.CreatedBy.ToUserShortResponse(),
		UpdatedBy:   note.UpdatedBy.ToUserShortResponse(),

		// Fields
		Subject: note.Subject,
		Note:    note.Note,
	}
}

func (note Notes) ToNoteResponses() []web.NoteResponse {
	noteResponses := []web.NoteResponse{}
	for _, note := range note {
		noteResponses = append(noteResponses, note.ToNoteResponse())
	}
	return noteResponses
}
