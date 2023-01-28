package web

import (
	"time"
)

type TestResponse struct {
	// Required Fields
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Required Fields
	Name         string     `json:"name"`
	Address      string     `json:"address"`
	Latitude     *float32   `json:"latitude"`
	Longitude    *float32   `json:"longitude"`
	Phone        *string    `json:"phone"`
	PlaceOfBirth *string    `json:"place_of_birth"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	Gender       string     `json:"gender"`
	Religion     string     `json:"religion"`
	Email        *string    `json:"email"`
	Ktp          *string    `json:"ktp"`
}
