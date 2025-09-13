package model

type RegistrationUserRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	StudentNisn string `json:"student_nisn" validate:"required"`
}
