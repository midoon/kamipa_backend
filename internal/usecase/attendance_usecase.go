package usecase

import (
	"context"
	"net/http"
	"time"

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

func (u *attendanceUsecase) GetAttendancesByStudentIdPaginated(
	ctx context.Context,
	userId string,
	page, size int,
) ([]model.AttendanceData, model.PageMetadata, error) {

	user, err := u.userRepository.GetById(ctx, userId)
	if err != nil {
		return []model.AttendanceData{}, model.PageMetadata{}, helper.NewCustomError(http.StatusInternalServerError, "Error get user data", err)
	}

	student, err := u.studentRepository.GetByNisn(ctx, user.StudentNisn)

	if err != nil {
		return []model.AttendanceData{}, model.PageMetadata{}, helper.NewCustomError(http.StatusInternalServerError, "Error get student data", err)
	}

	attendances, total, err := u.attendanceRepository.GetByStudentIdPaginated(ctx, student.ID, page, size)
	if err != nil {
		return []model.AttendanceData{}, model.PageMetadata{}, helper.NewCustomError(http.StatusInternalServerError, "Error get attendance data", err)
	}

	totalPages := int((total + int64(size) - 1) / int64(size))
	pageMetadata := model.PageMetadata{
		Page:       page,
		Size:       size,
		TotalPages: totalPages,
		TotalItems: int(total),
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

	return attendancesList, pageMetadata, nil
}

func (u *attendanceUsecase) GetAttendanceSummary(ctx context.Context, userId string) ([]model.AttendanceSummary, error) {

	user, err := u.userRepository.GetById(ctx, userId)
	if err != nil {
		return []model.AttendanceSummary{}, helper.NewCustomError(http.StatusInternalServerError, "Error get user data", err)
	}

	student, err := u.studentRepository.GetByNisn(ctx, user.StudentNisn)

	if err != nil {
		return []model.AttendanceSummary{}, helper.NewCustomError(http.StatusInternalServerError, "Error get student data", err)
	}

	attendances, err := u.attendanceRepository.GetByStudentId(ctx, student.ID)

	if err != nil {
		return []model.AttendanceSummary{}, helper.NewCustomError(http.StatusInternalServerError, "Error get attendance data", err)
	}

	atetndancesSummaryMap := make(map[string]*model.AttendanceSummary)
	for _, val := range attendances {
		summary, exists := atetndancesSummaryMap[val.Activity.Name]
		if !exists {
			summary = &model.AttendanceSummary{
				Activity:    val.Activity.Name,
				StartDate:   val.Date,
				CurrentDate: time.Now(),
				Hadir:       0,
				Izin:        0,
				Sakit:       0,
				Alpha:       0,
			}
			atetndancesSummaryMap[val.Activity.Name] = summary
		}

		switch val.Status {
		case "hadir":
			summary.Hadir++
		case "izin":
			summary.Izin++
		case "sakit":
			summary.Sakit++
		case "alpha":
			summary.Alpha++
		}
	}
	attendanceSummaries := []model.AttendanceSummary{}
	for _, summary := range atetndancesSummaryMap {
		attendanceSummaries = append(attendanceSummaries, *summary)
	}

	return attendanceSummaries, nil
}
