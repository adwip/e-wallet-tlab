package responses

type TopUpResp struct {
	Balance       float64 `json:"balance"`
	TransactionId string  `json:"transaction_id"`
	Status        string  `json:"status"`
}
