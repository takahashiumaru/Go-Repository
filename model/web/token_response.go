package web

type TokenResponse struct {
	// Fields
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
