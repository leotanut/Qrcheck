package models

//Login credential
type LoginResp struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Loginfail struct {
	Data string `json:"data"`
}
