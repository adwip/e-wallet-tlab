package entities

import "time"

type Wallet struct {
	ID            uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	SecureId      string    `gorm:"column:secure_id;not null;uniqueIndex;type:varchar(255)"`
	UserID        string    `gorm:"column:user_id;not null;type:varchar(255)"`
	Balance       float64   `gorm:"column:balance;not null;default:0;type:decimal(10,2)"`
	AccountNumber string    `gorm:"column:account_number;not null;uniqueIndex;type:varchar(255)"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;index"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime;index"`
}

func (Wallet) TableName() string {
	return "wallets"
}
