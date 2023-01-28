package domain

import (
	"time"

	"cek/model/web"

	"gorm.io/gorm"
)

type Tests []Test
type Test struct {
	// Required Fields
	gorm.Model
	ID          string    `gorm:"size:7;primaryKey"`
	CreatedAt   time.Time `gorm:""`
	CreatedByID uint      `gorm:""`
	UpdatedAt   time.Time `gorm:""`
	UpdatedByID uint      `gorm:""`
	DeletedByID *uint     `gorm:""`

	// Fields
	Name         string     `gorm:"size:500;not null"`
	Address      string     `gorm:"size:500;not null"`
	Latitude     *float32   `gorm:""`
	Longitude    *float32   `gorm:""`
	Phone        *string    `gorm:"size:15"`
	PlaceOfBirth *string    `gorm:"size:50"`
	DateOfBirth  *time.Time `gorm:"null"`
	Gender       string     `gorm:"size:20;not null"`
	Religion     string     `gorm:"size:25;not null"`
	Email        *string    `gorm:"size:50;unique;null"`
	Ktp          *string    `gorm:"size:25;unique;null"`
}

func (test *Test) ToTestResponse() web.TestResponse {

	return web.TestResponse{
		// Required Fields
		ID:        test.ID,
		CreatedAt: test.CreatedAt,
		UpdatedAt: test.UpdatedAt,

		// Fields
		Name:         test.Name,
		Address:      test.Address,
		Latitude:     test.Latitude,
		Longitude:    test.Longitude,
		Phone:        test.Phone,
		PlaceOfBirth: test.PlaceOfBirth,
		DateOfBirth:  test.DateOfBirth,
		Gender:       test.Gender,
		Religion:     test.Religion,
		Email:        test.Email,
		Ktp:          test.Ktp,
	}
}

func (tests Tests) ToTestResponses() []web.TestResponse {
	customerResponses := []web.TestResponse{}
	for _, customer := range tests {
		customerResponses = append(customerResponses, customer.ToTestResponse())
	}
	return customerResponses
}
