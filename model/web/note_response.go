package web

import (
	"time"
)

type NoteResponse struct {
	// Required Fields
	ID          uint              `json:"id"`
	CreatedByID uint              `json:"created_by_id"`
	UpdatedByID uint              `json:"updated_by_id"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	CreatedBy   UserShortResponse `json:"created_by"`
	UpdatedBy   UserShortResponse `json:"updated_by"`

	// Required Fields
	Subject string `json:"subject"`
	Note    string `json:"address"`
}
