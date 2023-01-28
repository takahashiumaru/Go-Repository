package web

type UserUpdateRequest struct {
	// Fields
	Name       string `json:"name" validate:"required,min=1,max=50"`
	Role       string `json:"role" validate:"required,min=1,max=25"`
	JoinDate   string `json:"join_date" validate:"required,period_day"`
	ResignDate string `json:"resign_date" validate:"max=8"`
	Email      string `json:"email" validate:"required,min=1,max=50"`
	Phone      string `json:"phone" validate:"required,min=1,max=20"`
}

type UserChangePasswordRequest struct {
	OldPassword    string `json:"old_password" validate:"required,min=6" `
	NewPassword    string `json:"new_password" validate:"required,min=6" `
	RetypePassword string `json:"retype_password" validate:"required,min=6" `
}
