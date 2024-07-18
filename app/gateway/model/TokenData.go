package model

type UserResp struct {
	UserId     string `json:"id"`
	Token      string `json:"token"`
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
