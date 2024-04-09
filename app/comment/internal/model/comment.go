package model

import "github.com/oigi/Magikarp/config/model"

type Comment struct {
	model.Model
	UserID  int64  `json:"user_id" gorm:"index;not null;comment:用户id;NOT NULL"`
	VideoID int64  `json:"video_id" gorm:"not null;comment:视频id;NOT NULL"`
	Content string `json:"content" gorm:"not null;comment:评论内容;NOT NULL"`
}
