package responses

type TransactionHistoryListResp struct {
	WalletId     string                   `json:"wallet_id"`
	Transactions []TransactionHistoryResp `json:"transactions"`
}

type TransactionHistoryResp struct {
	TransactionId   string  `json:"transaction_id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	TransactionDate string  `json:"transaction_date"`
	Type            string  `json:"type"`
	Description     string  `json:"description"`
}
