package dao

import (
	"context"
	publishModel "github.com/oigi/Magikarp/app/publish/internal/model"
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

func (f *PublishDao) InsertVideo(userId int64, title string, playUrl string, coverUrl string, label string, category string) (id int64, err error) {
	var video publishModel.Video
	video = publishModel.Video{
		AuthorId: userId,
		Title:    title,
		CoverUrl: coverUrl,
		Category: category,
		Label:    label,
		PlayUrl:  playUrl,
	}
	if err = f.Create(&video).Error; err != nil {
		return 0, errors.Wrap(err, "failed to create video")
	}
	return video.ID, nil
}
