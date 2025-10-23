package simipa_entity

import "time"

type Attendance struct {
	ID         int64     `gorm:"collumn:id;primaryKey"`
	StudentId  int64     `gorm:"collumn:student_id"`
	ActivityId int64     `gorm:"collumn:activity_id"`
	GroupId    int64     `gorm:"collumn:group_id"`
	Date       time.Time `gorm:"collumn:date"`
	Status     string    `gorm:"collumn:status"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *Attendance) TableName() string {
	return "attendances"
}
