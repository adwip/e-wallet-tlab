package dtos

import "time"

type TransactionHistoryDto struct {
	OperationId     string    `gorm:"operation_id"`
	Amount          float64   `gorm:"amount"`
	Status          string    `gorm:"status"`
	TransactionDate time.Time `gorm:"transaction_date"`
	Type            string    `gorm:"type"`
	Description     string    `gorm:"description"`
}
