package request

type Login struct {
	Username  string `json:"username"`  // 用户名
	Email     string `json:"mail"`      //邮箱
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Register User register structure
type Register struct {
	Password string `json:"passWord" example:"密码"`
	NickName string `json:"nickName" example:"昵称"`
	//HeaderImg string `json:"headerImg" example:"头像链接"`
	//AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable int `json:"enable" swaggertype:"string" example:"int 是否启用"`
	//Role   string `json:"role" example:"权限"`
	//AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	//Phone        string `json:"phone" example:"电话号码"`
	Email string `json:"email" example:"电子邮箱"`
}
