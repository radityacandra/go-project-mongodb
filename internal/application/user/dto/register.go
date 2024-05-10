package dto

type RegisterRequest struct {
	Username    string `json:"username" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
}
