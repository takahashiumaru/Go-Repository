package domain

import (
	"cek/model/web"

	"gorm.io/gorm"
)

type Users []User
type User struct {
	// Required Fields
	gorm.Model
	ID          uint   `gorm:"primarykey"`
	CreatedByID string `gorm:""`
	UpdatedByID string `gorm:""`
	DeletedByID string `gorm:""`

	Nip        string    `gorm:"size:20;not null;primarykey"`
	Password   string    `gorm:"not null"`
	Name       string    `gorm:"size:50;not null"`
	Role       string    `gorm:"size:30;not null"`
	JoinDate   string    `gorm:"size:8;not null"`
	ResignDate string    `gorm:"size:8"`
	Phone      string    `gorm:"size:15;not null"`
	Email      string    `gorm:"size:50;not null"`
	Image      string    `gorm:"size:500;"`
	Session    []Session `gorm:"foreignKey:UserID"`
}

func (user *User) ToUserResponse() web.UserResponse {
	return web.UserResponse{
		// Required Fields
		ID:          user.ID,
		CreatedByID: user.CreatedByID,
		UpdatedByID: user.UpdatedByID,

		// Fields
		Nip:        user.Nip,
		Name:       user.Name,
		Role:       user.Role,
		JoinDate:   user.JoinDate,
		ResignDate: user.ResignDate,
		Email:      user.Email,
		Phone:      user.Phone,
		Image:      user.Image,
	}
}

func (user *User) ToUserShortResponse() web.UserShortResponse {
	return web.UserShortResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (users Users) ToUserResponses() []web.UserResponse {
	userResponses := []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, user.ToUserResponse())
	}
	return userResponses
}
