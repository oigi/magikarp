package video

import "github.com/oigi/Magikarp/global"

type VStar struct {
	global.MODEL
	VideoID   uint `json:"video_id" gorm:"not null;comment:视频ID"`
	UserID    uint `json:"user_id" gorm:"comment:用户ID"`
	IsDeleted int  `json:"is_deleted" gorm:"default:0;comment:是否已删除"`
}
