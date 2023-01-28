package web

type UserCreateRequest struct {
	// Fields
	Name       string `json:"name" validate:"required,min=1,max=50"`
	Nip        string `json:"nip" validate:"required,min=1,max=25"`
	Password   string `json:"password" validate:"required,min=1,max=100"`
	Role       string `json:"role" validate:"required,min=1,max=25"`
	JoinDate   string `json:"join_date" validate:"required,period_day"`
	Email      string `json:"email" validate:"required,min=1,max=50"`
	Phone      string `json:"phone" validate:"required,min=1,max=20"`
}
