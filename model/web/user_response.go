package web

type UserResponse struct {
	// Fields
	ID          uint `json:"id"`
	CreatedByID string `json:"created_by_id"`
	UpdatedByID string `json:"updated_by_id"`
	Nip         string `json:"nip"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	JoinDate    string `json:"join_date"`
	ResignDate  string `json:"resign_date"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Image       string `json:"image"`
}
type UserShortResponse struct {
	// Fields
	ID   uint `json:"id"`
	Name string `json:"name"`
}
