package video

import "github.com/oigi/Magikarp/global"

type Category struct {
    global.MODEL
    Name        string `json:"name" gorm:"comment:类型名称"`
    Description string `json:"description" gorm:"comment:描述"`
    Open        int    `json:"open" gorm:"default:0;comment:是否开放"`
    Icon        string `json:"icon" gorm:"comment:图标"`
    Sort        int    `json:"sort" gorm:"default:0;comment:排序"`
    LabelNames  string `json:"label_names" gorm:"comment:标签名称"`
    IsDeleted   int    `json:"is_deleted" gorm:"default:0;comment:是否已删除"`
}
