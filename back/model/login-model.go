package model

type LoginModel struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}