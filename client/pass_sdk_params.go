package client

type createSDKTokenBody struct {
	User_id        string `json:"user_id"`
	Code_challenge string `json:"code_challenge"`
}
