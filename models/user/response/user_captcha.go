package response

type UserCaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	OpenCaptcha   bool   `json:"openCaptcha"`
}
