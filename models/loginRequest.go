package models

//Login credential
type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
