package dto

type AuthenticateRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiredAt   int64  `json:"expiredAt"`
}
