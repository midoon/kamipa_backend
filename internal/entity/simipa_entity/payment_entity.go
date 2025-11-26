package simipa_entity

import "time"

type Payment struct {
	ID          int64     `gorm:"collumn:id;primaryKey"`
	StudentId   int64     `gorm:"collumn:student_id"`
	FeeId       int64     `gorm:"collumn:fee_id"`
	Amount      float64   `gorm:"collumn:amount"`
	PaymentDate time.Time `gorm:"collumn:payment_date"`
	Description string    `gorm:"collumn:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *Payment) TableName() string {
	return "payments"
}
