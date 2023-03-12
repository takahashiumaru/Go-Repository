package web

type NoteUpdateRequest struct {
	// Required Fields
	Subject string `json:"subject" validate:"required:min=1"`
	Note    string `json:"note" validate:"required:min=1"`
}
