package web

type NoteCreateRequest struct {
	// Required Fields
	Subject string `json:"subject" validate:"required"`
	Note    string `json:"note" validate:"required"`
}
