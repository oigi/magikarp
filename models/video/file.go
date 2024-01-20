package video

import (
    "github.com/oigi/Magikarp/global"
)

type File struct {
    global.MODEL
    FileKey   string `json:"file_key" gorm:"comment:文件键"`
    Format    string `json:"format" gorm:"comment:文件格式"`
    Type      string `json:"type" gorm:"comment:文件类型"`
    Duration  string `json:"duration" gorm:"comment:文件时长"`
    Size      int64  `json:"size" gorm:"comment:文件大小"`
    UserID    int    `json:"user_id" gorm:"index;comment:用户ID"`
    IsDeleted int    `json:"is_deleted" gorm:"comment:是否已删除"`
}
