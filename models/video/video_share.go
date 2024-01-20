package video

import "github.com/oigi/Magikarp/global"

type VShare struct {
	global.MODEL
	VideoID uint   `json:"video_id" gorm:"comment:视频ID"`
	IP      string `json:"ip" gorm:"comment:IP地址"`
	UserID  uint   `json:"user_id" gorm:"comment:用户ID"`
}
