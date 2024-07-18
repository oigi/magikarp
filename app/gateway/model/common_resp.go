package model

type CommonResp struct {
	StatusCode int    `form:"status_code" json:"status_code" xml:"status_code"`
	StatusMsg  string `form:"status_msg" json:"status_msg" xml:"status_msg"`
}
