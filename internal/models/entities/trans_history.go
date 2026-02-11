package entities

import "time"

type TransactionHistories struct {
	ID            uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	SecureId      string    `gorm:"column:secure_id;not null;uniqueIndex;type:varchar(255)"`
	TransactionId string    `gorm:"column:transaction_id;not null;uniqueIndex;type:varchar(255)"`
	Status        string    `gorm:"column:status;not null;type:varchar(255)"`
	WalletId      string    `gorm:"column:wallet_id;not null;type:varchar(255)"`
	Amount        float64   `gorm:"column:amount;not null;type:decimal(10,2)"`
	Type          string    `gorm:"column:type;not null;type:varchar(255)"`
	Description   string    `gorm:"column:description;not null;type:varchar(255)"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;index"`
}

func (TransactionHistories) TableName() string {
	return "transactions_histories"
}
