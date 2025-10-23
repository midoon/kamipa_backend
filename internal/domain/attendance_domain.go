package domain

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/model"
)

type AttendanceRepository interface {
	GetByStudentId(ctx context.Context, studentId int64) ([]simipa_entity.Attendance, error)
}

type AttendanceUsecase interface {
	GetAttendances(ctx context.Context, userId string) ([]model.AttendanceData, error)
}
