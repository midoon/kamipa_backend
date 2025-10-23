package controller

import (
	"net/http"

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
