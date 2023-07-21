package models

type Login struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}
