package model

type RegistrationUserRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	StudentNisn string `json:"student_nisn" validate:"required"`
}

type LoginUserRequest struct {
	StudentNisn string `json:"student_nisn" validate:"required"`
	Password    string `json:"password" validate:"required,min=6"`
}

type TokenDataResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
