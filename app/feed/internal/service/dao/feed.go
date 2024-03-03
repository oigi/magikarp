package dao

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/initialize/mysql"
	feedModel "github.com/oigi/Magikarp/models/feed"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FeedDao struct {
	*gorm.DB
}

func NewFeedDao(ctx context.Context) *FeedDao {
	return &FeedDao{mysql.NewDBClient(ctx)}
}

// 获取视频信息

// 获取视频播放量排行榜

// 获取对应分类的视频

// 获取关注人的视频

// CreateUser 创建视频
func (dao *FeedDao) CreateUser(req *feed.CreateVideoReq) (err error) {
	var video feedModel.Videos
	video = feedModel.Videos{
		AuthorId: int(req.ActorId),
		Title:    req.Title,
		CoverUrl: req.CoverUrl,
		Category: int(req.Category),
		//  TODO
	}
	if err = dao.Create(&video).Error; err != nil {
		return errors.Wrap(err, "failed to create video")
	}
	return err
}

// 更新视频
// 删除视频
// 定时同步视频播放量 点赞量 收藏数
// 生成视频推荐
