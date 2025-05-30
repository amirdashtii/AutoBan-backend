package dto

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
    PhoneNumber string `validate:"required,iranphone" json:"phone_number"`
    Password    string `validate:"required,min=8,password" json:"password"`
}

type LoginRequest struct {
	PhoneNumber string `validate:"required,iranphone" json:"phone_number"`
	Password    string `validate:"required,min=8,password" json:"password"`
}

type LoginResponse struct {
	Token Token `json:"token"`
}
type LogoutRequest struct {
	RefreshToken string `validate:"required" json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `validate:"required" json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Token Token `json:"token"`
}
