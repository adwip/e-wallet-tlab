package dtos

import "time"

type TransactionHistoryDto struct {
	TransactionId   string    `gorm:"transaction_id"`
	Amount          float64   `gorm:"amount"`
	Status          string    `gorm:"status"`
	TransactionDate time.Time `gorm:"transaction_date"`
	Type            string    `gorm:"type"`
	Description     string    `gorm:"description"`
}
