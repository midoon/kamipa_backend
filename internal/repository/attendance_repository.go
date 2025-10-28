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

func (r *attendanceRepository) GetByStudentIdPaginated(ctx context.Context, studentId int64, page int, size int) ([]simipa_entity.Attendance, int64, error) {

	var attendances []simipa_entity.Attendance
	var total int64

	if err := r.simipaDB.WithContext(ctx).Model(&simipa_entity.Attendance{}).Where("student_id = ?", studentId).Count(&total).Error; err != nil {
		return []simipa_entity.Attendance{}, 0, err
	}

	offset := (page - 1) * size // offset adalah dimulai dari data ke berapa data tsb diambil
	if err := r.simipaDB.WithContext(ctx).Preload("Activity").Where("student_id = ?", studentId).Order("date DESC").Limit(size).Offset(offset).Find(&attendances).Error; err != nil {
		return []simipa_entity.Attendance{}, 0, err
	}

	return attendances, total, nil
}
