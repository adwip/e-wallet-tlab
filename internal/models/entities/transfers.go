package entities

import "time"

type Transfers struct {
	ID             uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	SecureId       string    `gorm:"column:secure_id;not null;uniqueIndex;type:varchar(255)"`
	Amount         float64   `gorm:"column:amount;not null;type:decimal(10,2)"`
	SenderId       string    `gorm:"column:sender_id;not null;type:varchar(255)"`
	WalletSourceId string    `gorm:"column:wallet_source_id;not null;type:varchar(255)"`
	ReceiverId     string    `gorm:"column:receiver_id;not null;type:varchar(255)"`
	WalletDestId   string    `gorm:"column:wallet_dest_id;not null;type:varchar(255)"`
	Status         string    `gorm:"column:status;not null;type:varchar(255)"`
	Note           string    `gorm:"column:note;not null;type:varchar(255)"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime;index"`
}

func (Transfers) TableName() string {
	return "transfers"
}
