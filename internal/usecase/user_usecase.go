package usecase

import "github.com/midoon/kamipa_backend/internal/domain"

type userUsecase struct {
	userRepo    domain.UserRepository
	studentRepo domain.StudentRepository
}

func NewUserRepository(userRepo domain.UserRepository, studentRepo domain.StudentRepository) domain.UserUseCase {
	return &userUsecase{
		userRepo:    userRepo,
		studentRepo: studentRepo,
	}
}
