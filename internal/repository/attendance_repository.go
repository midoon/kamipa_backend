package repository

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"gorm.io/gorm"
)

type attendanceRepository struct {
	simipaDB *gorm.DB
}

func NewAttendanceRepository(simipaDB *gorm.DB) domain.AttendanceRepository {
	return &attendanceRepository{
		simipaDB: simipaDB,
	}
}

func (r *attendanceRepository) GetByStudentId(ctx context.Context, studentId int64) ([]simipa_entity.Attendance, error) {
	var attendances []simipa_entity.Attendance

	err := r.simipaDB.WithContext(ctx).Preload("Activity").Where("student_id = ?", studentId).Find(&attendances).Error

	if err != nil {
		return []simipa_entity.Attendance{}, err
	}
	return attendances, nil
}
