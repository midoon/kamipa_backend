package domain

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/model"
)

type AttendanceRepository interface {
	GetByStudentId(ctx context.Context, studentId int64) ([]simipa_entity.Attendance, error)
	GetByStudentIdPaginated(ctx context.Context, studentId int64, page int, size int) ([]simipa_entity.Attendance, int64, error)
}

type AttendanceUsecase interface {
	GetAttendances(ctx context.Context, userId string) ([]model.AttendanceData, error)
	GetAttendancesByStudentIdPaginated(
		ctx context.Context,
		userId string,
		page, size int,
	) ([]model.AttendanceData, model.PageMetadata, error)
	GetAttendanceSummary(ctx context.Context, userId string) ([]model.AttendanceSummary, error)
}
