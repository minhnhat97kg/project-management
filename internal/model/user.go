package model

type RegisterRequest struct {
	UserName string `validate:"required,min=3,max=40" json:'userName'`
	Password string `validate:"required,min=6" json:'password' `
}
