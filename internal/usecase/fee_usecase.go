package usecase

import (
	"context"
	"net/http"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type feeUsecase struct {
	feeRepository     domain.FeeRepository
	userRepository    domain.UserRepository
	studentRepository domain.StudentRepository
}

func NewFeeUsecase(feeRepository domain.FeeRepository, userRepository domain.UserRepository, studentRepository domain.StudentRepository) domain.FeeUsecase {
	return &feeUsecase{
		feeRepository:     feeRepository,
		userRepository:    userRepository,
		studentRepository: studentRepository,
	}
}

func (u *feeUsecase) GetFees(ctx context.Context, userId string) ([]model.FeeList, error) {
	user, err := u.userRepository.GetById(ctx, userId)
	if err != nil {
		return []model.FeeList{}, helper.NewCustomError(http.StatusInternalServerError, "Error get user data", err)
	}

	student, err := u.studentRepository.GetByNisn(ctx, user.StudentNisn)

	if err != nil {
		return []model.FeeList{}, helper.NewCustomError(http.StatusInternalServerError, "Error get student data", err)
	}

	fees, err := u.feeRepository.GetByStudentId(ctx, student.ID)
	if err != nil {
		return []model.FeeList{}, helper.NewCustomError(http.StatusInternalServerError, "Error get fee data", err)
	}

	feeList := []model.FeeList{}

	for _, val := range fees {
		feeData := model.FeeList{
			ID:     val.ID,
			Name:   val.PaymentType.Name,
			Status: val.Status,
		}

		feeList = append(feeList, feeData)
	}

	return feeList, nil
}

func (u *feeUsecase) GetFeeDetail(ctx context.Context, feeId int64) (model.FeeDetail, error) {
	fee, err := u.feeRepository.GetByFeeId(ctx, feeId)
	if err != nil {
		return model.FeeDetail{}, helper.NewCustomError(http.StatusInternalServerError, "Error get fee data", err)
	}
	feeDetail := model.FeeDetail{
		ID:              fee.ID,
		Name:            fee.PaymentType.Name,
		Amount:          fee.Amount,
		Status:          fee.Status,
		RemainingAmount: fee.Amount - fee.PaidAmount,
		DueDate:         fee.DueDate,
	}
	return feeDetail, nil
}
