package kamipa_entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ErrorPaymentLog struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	OrderID   string    `gorm:"size:100;index" json:"order_id"`
	FeeID     int64     `gorm:"size:100" json:"fee_id"`
	Raw       string    `gorm:"collumn:raw;type:text" json:"raw"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *ErrorPaymentLog) TableName() string {
	return "error_payment_logs"
}

func (u *ErrorPaymentLog) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		id := uuid.New().String()
		u.ID = id
	}

	return nil
}
