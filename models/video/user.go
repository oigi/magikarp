package video

import (
    "github.com/oigi/Magikarp/global"
)

type User struct {
    global.MODEL
    UID       int    `json:"uid" gorm:"index;comment:用户UID"`                           // 用户UUID
    Username  string `json:"userName" gorm:"index;comment:用户登录名"`                   // 用户登录名
    Password  string `json:"-"  gorm:"comment:用户登录密码"`                             // 用户登录密码
    NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`          // 用户昵称
    HeaderImg string `json:"headerImg" gorm:";comment:用户头像"`                         // 用户头像
    Email     string `json:"email"  gorm:"comment:用户邮箱"`                             // 用户邮箱
    Enable    int    `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}
