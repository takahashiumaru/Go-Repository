package web

type RefreshTokenCreateRequest struct {
	// Fields
	RefreshToken string `json:"refresh_token" validate:"required,jwt"`
}
