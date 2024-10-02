package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
