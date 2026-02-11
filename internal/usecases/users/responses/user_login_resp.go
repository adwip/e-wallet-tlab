package responses

type UserLoginResp struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
