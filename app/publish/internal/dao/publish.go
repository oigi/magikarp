package dao

import (
	"context"
	publishModel "github.com/oigi/Magikarp/app/publish/internal/model"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"github.com/oigi/Magikarp/pkg/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PublishDao struct {
	*gorm.DB
}

func NewPublishDao(ctx context.Context) *PublishDao {
	return &PublishDao{
		mysql.NewDBClient(ctx)}
}

func (f *PublishDao) InsertVideo(req *publish.CreateVideoRequest) (err error) {
	var video publishModel.Videos
	video = publishModel.Videos{
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

// FindVideoListByUserId 获取用户的全部视频
func (f *PublishDao) FindVideoListByUserId(req *publish.ListVideoRequest) (videos []publishModel.Videos, err error) {
	if err := f.Where("actor_id = ?", req.ActorId).First(&videos).Error; err != nil {
		err = errors.Wrap(err, "查询视频失败")
	}
	return
}
