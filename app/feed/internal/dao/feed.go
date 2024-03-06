package dao

import (
    "context"
    feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
    "github.com/oigi/Magikarp/grpc/pb/feed"
    "github.com/oigi/Magikarp/pkg/mysql"
    "github.com/pkg/errors"
    "gorm.io/gorm"
)

type FeedDao struct {
    *gorm.DB
}

func NewFeedDao(ctx context.Context) *FeedDao {
    return &FeedDao{
        mysql.NewDBClient(ctx)}
}

// FindAllVideos 获取全部视频
func (f *FeedDao) FindAllVideos() (videos []feedModel.Videos, err error) {
    if err := f.Find(&videos).Error; err != nil {
        err = errors.Wrap(err, "查询视频失败")
    }
    return
}

// FindVideoById 根据视频id获取视频信息
func (f *FeedDao) FindVideoById(req *feed.SearchVideoReq) (video feedModel.Videos, err error) {
    if err := f.Where("id = ?", req.VideoId).First(&video).Error; err != nil {
        err = errors.Wrap(err, "查询视频失败")
    }
    return
}

// FindVideoListByCategory 根据对应属性找视频
func (f *FeedDao) FindVideoListByCategory(req *feed.SearchVideoReq) (videos []feedModel.Videos, err error) {
    if err = f.Where("category = ?", req.Category).Find(&videos).Error; err != nil {
        err = errors.Wrap(err, "查询视频失败")
    }
    return
}

// FindVideoListByTable 根据对应标签找视频
func (f *FeedDao) FindVideoListByTable(req *feed.SearchVideoReq) (videos []feedModel.Videos, err error) {
    if err = f.Where("table = ?", req.Table).Find(&videos).Error; err != nil {
        err = errors.Wrap(err, "查询视频失败")
    }
    return
}

// FindVideoListByUserId 获取用户的全部视频
func (f *FeedDao) FindVideoListByUserId(req *feed.SearchVideoReq) (videos []feedModel.Videos, err error) {
    if err := f.Where("actor_id = ?", req.ActorId).First(&videos).Error; err != nil {
        err = errors.Wrap(err, "查询视频失败")
    }
    return
}

// InsertVideo 创建视频
func (f *FeedDao) InsertVideo(req *feed.CreateVideoReq) (err error) {
    var video feedModel.Videos
    video = feedModel.Videos{
        AuthorId: req.ActorId,
        Title:    req.Title,
        CoverUrl: req.CoverUrl,
        Category: req.Category,
        Label:    req.Label,
    }
    if err = f.Create(&video).Error; err != nil {
        return errors.Wrap(err, "failed to create video")
    }
    return
}

// UpdateVideoById 更新视频
func (f *FeedDao) UpdateVideoById(req *feed.DeleteVideoReq) (err error) {
    var video feedModel.Videos
    if err = f.Model(&video).Update("id = ?", req.VideoId).Error; err != nil {
        err = errors.Wrap(err, "")
    }
    return
}

// DeleteVideoById 删除视频
func (f *FeedDao) DeleteVideoById(req *feed.DeleteVideoReq) (err error) {
    var video feedModel.Videos
    if err = f.Model(&video).Delete("id = ?", req.VideoId).Error; err != nil {
        err = errors.Wrap(err, "")
    }
    return
}
