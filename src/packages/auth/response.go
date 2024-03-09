package auth

type TokenResponse struct {
	AccessToken  string `json:"accessToken" example:"string"`
	RefreshToken string `json:"refreshToken" example:"string"`
}
