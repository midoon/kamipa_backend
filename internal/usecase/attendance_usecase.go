package usecase

import (
	"context"
	"net/http"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type attendanceUsecase struct {
	attendanceRepository domain.AttendanceRepository
	userRepository       domain.UserRepository
	studentRepository    domain.StudentRepository
}

func NewAttendanceUsecase(attendanceRepository domain.AttendanceRepository, userRepository domain.UserRepository, studentRepository domain.StudentRepository) domain.AttendanceUsecase {
	return &attendanceUsecase{
		attendanceRepository: attendanceRepository,
		studentRepository:    studentRepository,
		userRepository:       userRepository,
	}
}

func (u *attendanceUsecase) GetAttendances(ctx context.Context, userId string) ([]model.AttendanceData, error) {

	user, err := u.userRepository.GetById(ctx, userId)
	if err != nil {
		return []model.AttendanceData{}, helper.NewCustomError(http.StatusInternalServerError, "Error get user data", err)
	}

	student, err := u.studentRepository.GetByNisn(ctx, user.StudentNisn)

	if err != nil {
		return []model.AttendanceData{}, helper.NewCustomError(http.StatusInternalServerError, "Error get student data", err)
	}

	attendances, err := u.attendanceRepository.GetByStudentId(ctx, student.ID)

	if err != nil {
		return []model.AttendanceData{}, helper.NewCustomError(http.StatusInternalServerError, "Error get attendance data", err)
	}

	attendancesList := []model.AttendanceData{}

	for _, val := range attendances {
		attendanceData := model.AttendanceData{
			StudentId: val.StudentId,
			Date:      val.Date,
			Status:    val.Status,
			Activity:  val.Activity.Name,
		}

		attendancesList = append(attendancesList, attendanceData)
	}

	return attendancesList, nil
}
