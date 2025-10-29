package controller

import (
	"net/http"
	"strconv"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type AttendanceController struct {
	attendanceUsecase domain.AttendanceUsecase
}

func NewAttendanceController(attendanceUsecase domain.AttendanceUsecase) *AttendanceController {
	return &AttendanceController{
		attendanceUsecase: attendanceUsecase,
	}
}

func (c *AttendanceController) GetAttendances(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Context().Value(helper.UserIDKey).(string)

	attendances, err := c.attendanceUsecase.GetAttendances(ctx, userId)
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	resp := model.ArrayResponse[model.AttendanceData]{
		Status:  true,
		Message: "Success get attendance data",
		Data:    attendances,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}

func (c *AttendanceController) GetAttendancesPaginated(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Context().Value(helper.UserIDKey).(string)
	page := getIntQuery(r, "page", 1)
	size := getIntQuery(r, "size", 10)

	attendances, metaData, err := c.attendanceUsecase.GetAttendancesByStudentIdPaginated(ctx, userId, page, size)
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	resp := model.PageResponse[model.AttendanceData]{
		Status:   true,
		Message:  "Success get attendance data",
		Data:     attendances,
		Metadata: metaData,
	}

	helper.WriteJSON(w, http.StatusOK, resp)

}

func getIntQuery(r *http.Request, key string, defaultVal int) int {
	valStr := r.URL.Query().Get(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

func (c *AttendanceController) GetAttendanceSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Context().Value(helper.UserIDKey).(string)
	attendanceSummaries, err := c.attendanceUsecase.GetAttendanceSummary(ctx, userId)
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	resp := model.ArrayResponse[model.AttendanceSummary]{
		Status:  true,
		Message: "Success get attendance summary data",
		Data:    attendanceSummaries,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}
