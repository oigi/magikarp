package model

import (
	"github.com/oigi/Magikarp/config/model"
)

type Relation struct {
	model.Model
	UserId         int64 `gorm:"column:user_id;not null;comment:用户id"`
	FollowedUserId int64 `gorm:"column:following_id;not null;comment:被关注用户id"`
	Status         int   `gorm:"column:status;not null;comment:关系状态"` // 1关注 2取消关注
}
