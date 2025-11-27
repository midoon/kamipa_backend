package simipa_entity

import "time"

type Fee struct {
	ID            int64     `gorm:"collumn:id;primaryKey"`
	StudentId     int64     `gorm:"collumn:student_id"`
	GradeFeeId    int64     `gorm:"collumn:grade_fee_id"`
	PaymentTypeId int64     `gorm:"collumn:payment_type_id"`
	Amount        float64   `gorm:"collumn:amount"`
	DueDate       time.Time `gorm:"collumn:due_date"`
	Status        string    `gorm:"collumn:status"`
	PaidAmount    float64   `gorm:"collumn:paid_amount"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	PaymentType PaymentType `gorm:"foreignKey:PaymentTypeId;references:ID"`
}

func (u *Fee) TableName() string {
	return "fees"
}
