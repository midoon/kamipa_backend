package kamipa_entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TopupLog struct {
	ID        string                 `gorm:"primaryKey" json:"id"`
	OrderID   string                 `gorm:"size:100;index" json:"order_id"`
	Event     string                 `gorm:"size:100" json:"event"`
	Status    string                 `gorm:"size:50" json:"status"`
	Raw       map[string]interface{} `gorm:"type:jsonb" json:"raw"`
	CreatedAt time.Time              `json:"created_at"`
}

func (u *TopupLog) TableName() string {
	return "topup_logs"
}

func (u *TopupLog) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		id := uuid.New().String()
		u.ID = id
	}

	return nil
}
