package responses

type TransferResp struct {
	TransactionId string `json:"transaction_id"`
	Status        string `json:"status"`
}
