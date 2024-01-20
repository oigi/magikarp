package user

import "github.com/oigi/Magikarp/global"

type UFavorites struct {
	global.MODEL
	Name        string `json:"name" gorm:"comment:收藏夹名称"`
	Description string `json:"description" gorm:"comment:收藏夹描述"`
	UserID      int    `json:"user_id" gorm:"index;comment:用户ID"`
	IsDeleted   int    `json:"is_deleted" gorm:"default:0;comment:是否已删除"`
}
