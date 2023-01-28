package web

type UserResponse struct {
	// Fields
	ID         uint   `json:"id"`
	Nip        string `json:"nip"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	JoinDate   string `json:"join_date"`
	ResignDate string `json:"resign_date"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}
type UserShortResponse struct {
	// Fields
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
