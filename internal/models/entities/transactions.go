package entities

import "time"

type Transaction struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	SecureId    string    `gorm:"column:secure_id;not null;uniqueIndex;type:varchar(255)"`
	WalletID    string    `gorm:"column:wallet_id;not null;type:varchar(255)"`
	Amount      float64   `gorm:"column:amount;not null;type:decimal(10,2)"`
	ActionType  string    `gorm:"column:action_type;not null;type:varchar(255)"`
	Note        string    `gorm:"column:note;not null;type:varchar(255)"`
	OperationId string    `gorm:"column:operation_id;not null;type:varchar(255)"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;index"`
}

func (Transaction) TableName() string {
	return "transactions"
}
