package model

import (
	"database/sql"
	"github.com/oigi/Magikarp/config/model"
)

type Videos struct {
	model.Model
	AuthorId      int64          `json:"author_id" gorm:"column:author_id;type:int(11) unsigned;comment:上传用户Id;NOT NULL"`
	Title         string         `json:"title" gorm:"column:title;type:varchar(255);comment:视频标题;NOT NULL"`
	CoverUrl      string         `json:"cover_url" gorm:"column:cover_url;type:varchar(255);comment:封面url;NOT NULL"`
	PlayUrl       string         `json:"play_url" gorm:"column:play_url;type:varchar(255);comment:视频播放url;NOT NULL"`
	PlayCount     int64          `json:"play_count" gorm:"column:play_count;comment:视频播放量;NOT NULL"`
	FavoriteCount int64          `json:"favorite_count" gorm:"column:favorite_count;type:int(11) unsigned;default:0;comment:点赞数;NOT NULL"`
	StarCount     int64          `json:"star_count" gorm:"column:star_count;type:int(11);comment:收藏数;NOT NULL"`
	CommentCount  int64          `json:"comment_count" gorm:"column:comment_count;type:int(11) unsigned;default:0;comment:评论数目;NOT NULL"`
	Category      string         `json:"category" gorm:"column:category;type:int(11);comment:视频分类;NOT NULL"`
	Duration      sql.NullString `json:"duration" gorm:"column:duration;type:varchar(255);comment:视频时长"`
	Label         string         `json:"label" gorm:"column:label;type:varchar(255);comment:视频标签;NOT NULL"`
	Open          bool           `json:"open" gorm:"column:open;default:'true';comment:是否公开"`
}
