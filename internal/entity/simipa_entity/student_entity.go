package simipa_entity

import "time"

type Student struct {
	ID        int64     `gorm:"collumn:id;primaryKey"`
	GroupId   int64     `gorm:"collumn:group_id"`
	Name      string    `gorm:"collumn:name"`
	Nisn      string    `gorm:"collumn:nisn"`
	Gender    string    `gorm:"collumn:gender"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *Student) TableName() string {
	return "students"
}
