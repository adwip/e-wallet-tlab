package responses

type UserRegistrationResp struct {
	Name     string `json:"name"`
	SecureId string `json:"secure_id"`
}
