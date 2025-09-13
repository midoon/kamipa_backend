package kamipa_entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string    `gorm:"columnn:id;primaryKey"`
	StudentNisn string    `gorm:"column:student_nisn"`
	Email       string    `gorm:"column:email"`
	Password    string    `gorm:"column:password"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		id := uuid.New().String()
		u.ID = id
	}

	return nil
}
