package video

import (
	"gorm.io/gorm"
)

// Video 视频表
type Video struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null;comment:标题"`   // 标题
	Description string `json:"description" gorm:"comment:描述"`      // 描述
	URL         string `json:"url" gorm:"comment:URL"`             // URL
	UserID      int64  `json:"user_id" gorm:"comment:用户ID"`        // 用户ID
	TypeID      int64  `json:"type_id" gorm:"comment:类型ID"`        // 类型ID
	Open        int    `json:"open" gorm:"default:0;comment:是否开放"` // 是否开放 0开放1关闭
	Cover       string `json:"cover" gorm:"comment:封面"`            // 封面
	//AuditStatus      int    `json:"audit_status" gorm:"not null;comment:审计状态"`          // 审计状态
	Msg string `json:"msg" gorm:"comment:消息"` // 消息
	//AuditQueueStatus int    `json:"audit_queue_status" gorm:"default:0;comment:审计队列状态"` // 审计队列状态
	StartCount   int64 `json:"start_count" gorm:"default:0;comment:开始计数"`   // 开始计数
	ShareCount   int64 `json:"share_count" gorm:"default:0;comment:分享计数"`   // 分享计数
	HistoryCount int64 `json:"history_count" gorm:"default:0;comment:历史计数"` // 历史计数
}

// VideoType 视频分类表
type VideoType struct {
	gorm.Model
	Name        string `json:"name" gorm:"comment:名称"`             // 名称
	Description string `json:"description" gorm:"comment:描述"`      // 描述
	Open        int    `json:"open" gorm:"default:0;comment:是否开放"` // 是否开放
	Icon        string `json:"icon" gorm:"comment:图标"`             // 图标
	Sort        int    `json:"sort" gorm:"default:0;comment:排序"`   // 排序
	LabelNames  string `json:"label_names" gorm:"comment:标签名称"`    // 标签名称
}
