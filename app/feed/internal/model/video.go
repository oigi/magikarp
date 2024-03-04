package feed

import (
	"database/sql"
	"github.com/oigi/Magikarp/consts"
)

type Videos struct {
	AuthorId      int            `json:"author_id" gorm:"column:author_id;type:int(11) unsigned;comment:上传用户Id;NOT NULL"`
	Title         string         `json:"title" gorm:"column:title;type:varchar(255);comment:视频标题;NOT NULL"`
	CoverUrl      string         `json:"cover_url" gorm:"column:cover_url;type:varchar(255);comment:封面url;NOT NULL"`
	PlayUrl       string         `json:"play_url" gorm:"column:play_url;type:varchar(255);comment:视频播放url;NOT NULL"`
	PlaysCount    int            `json:"plays_count" gorm:"column:plays_count;comment:视频播放量;NOT NULL"`
	FavoriteCount int            `json:"favorite_count" gorm:"column:favorite_count;type:int(11) unsigned;default:0;comment:点赞数;NOT NULL"`
	StarCount     int            `json:"star_count" gorm:"column:star_count;type:int(11);comment:收藏数;NOT NULL"`
	CommentCount  int            `json:"comment_count" gorm:"column:comment_count;type:int(11) unsigned;default:0;comment:评论数目;NOT NULL"`
	Category      int            `json:"category" gorm:"column:category;type:int(11);comment:视频分类;NOT NULL"`
	Duration      sql.NullString `json:"duration" gorm:"column:duration;type:varchar(255);comment:视频时长"`
	Label         string         `json:"label" gorm:"column:label;type:varchar(255);comment:视频标签;NOT NULL"`
	Open          bool           `json:"open" gorm:"column:open;default:'true';comment:是否公开"`
	consts.Model
}
