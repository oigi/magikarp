package request

type Login struct {
	Username string `json:"username"` // 用户名
	Email    string `json:"mail"`     //邮箱
	Password string `json:"password"` // 密码
	//Captcha   string `json:"captcha"`   // 验证码
	//CaptchaId string `json:"captchaId"` // 验证码ID
}
