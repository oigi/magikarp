package dao

import (
    "context"
    feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
    "github.com/oigi/Magikarp/grpc/pb/feed"
    "github.com/oigi/Magikarp/initialize/mysql"
    "github.com/pkg/errors"
    "gorm.io/gorm"
    "time"
)

type FeedDao struct {
    db    *gorm.DB
    video feedModel.Videos
}

func NewFeedDao(ctx context.Context) *FeedDao {
    return &FeedDao{
        db: mysql.NewDBClient(ctx)}
}

// 获取视频信息

// 获取视频播放量排行榜

// GetVideoList 获取对应分类的视频
func (f *FeedDao) GetVideoList(req *feed.GetVideoListReq) (r feedModel.Videos, err error) {
    now := time.Now()
    last := now.AddDate(-1, 0, 0)
    if err = f.db.Limit(20).Where("created_at >= ?", last).Order("rand()").Find(&r).Error; err != nil {
        err = errors.Wrap(err, "查询视频失败")
    }
    return
}

// 获取关注人的视频

// CreateVideo 创建视频
func (f *FeedDao) CreateVideo(req *feed.CreateVideoReq) (err error) {
    f.video = feedModel.Videos{
        AuthorId: req.ActorId,
        Title:    req.Title,
        CoverUrl: req.CoverUrl,
        Category: req.Category,
        Label:    req.Label,
    }
    if err = f.db.Create(&f.video).Error; err != nil {
        return errors.Wrap(err, "failed to create video")
    }
    return
}

// UpdateVideo 更新视频
func (f *FeedDao) UpdateVideo(req *feed.DeleteVideoReq) (err error) {
    if err = f.db.Model(&f.video).Update("title = ?", req.VideoId).Error; err != nil {
        return errors.Wrap(err, "")
    }
    return err
}

// DeleteVideo 删除视频
func (f *FeedDao) DeleteVideo(req *feed.DeleteVideoReq) (err error) {
    if err = f.db.Model(&f.video).Delete("title = ?", req.VideoId).Error; err != nil {
        return errors.Wrap(err, "")
    }
    return err
}

// InitVideoList 未登陆初始化推送视频
func (f *FeedDao) InitVideoList() (r []feedModel.Videos, err error) {
    now := time.Now()
    last := now.AddDate(-1, 0, 0)
    if err = f.db.Limit(20).Where("created_at >= ?", last).Order("rand()").Find(&r).Error; err != nil {
        err = errors.Wrap(err, "查询视频失败")
    }
    return
}
