package kamipa_entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Topup struct {
	ID              string                 `gorm:"primaryKey" json:"id"`
	OrderID         string                 `gorm:"uniqueIndex;size:100" json:"order_id"`
	FeeID           int64                  `gorm:"size:100" json:"fee_id"`
	UserID          string                 `json:"user_id"`
	Amount          int64                  `json:"amount"`
	SnapToken       string                 `gorm:"size:255" json:"snap_token"`
	SnapTokenExpiry *time.Time             `json:"snap_token_expiry"`
	Status          string                 `gorm:"size:50;default:'pending'" json:"status"`
	PaymentType     string                 `gorm:"size:50" json:"payment_type"`
	PaymentCode     string                 `gorm:"size:255" json:"payment_code"`
	TransactionTime *time.Time             `json:"transaction_time"`
	SettlementTime  *time.Time             `json:"settlement_time"`
	RawResponse     map[string]interface{} `gorm:"type:jsonb" json:"raw_response"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

func (u *Topup) TableName() string {
	return "topups"
}

func (u *Topup) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		id := uuid.New().String()
		u.ID = id
	}

	return nil
}
