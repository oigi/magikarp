package dao

import (
	"context"
	favoriteModel "github.com/oigi/Magikarp/app/favorite/internal/model"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"github.com/oigi/Magikarp/pkg/mysql"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{
		mysql.NewDBClient(ctx)}
}

func (f *FavoriteDao) QueryFavoriteCount(req *favorite.FavoriteCountReq) (resp *favorite.FavoriteCountResp, err error) {
	var favoriteModel favoriteModel.Favorite
	resp = &favorite.FavoriteCountResp{
		VideoFavoriteCount: make(map[int64]int64),
	}

	for _, videoID := range req.VideoIdList {
		var count int64

		if err := f.Model(&favoriteModel).Where("video_id = ?", videoID).Count(&count).Error; err != nil {
			return nil, err
		}

		resp.VideoFavoriteCount[videoID] = count
	}

	return resp, nil
}

// UpdateFavoriteCount 定时更新
func (f *FavoriteDao) UpdateFavoriteCount(req *favorite.FavoriteCountReq) (resp map[int64]int64, err error) {
	resp = make(map[int64]int64)
	var favoriteModel favoriteModel.Favorite
	counts, err := QueryFavoriteCount(req.VideoIdList)
	if err != nil {
		return nil, err
	}
	for videoID, count := range counts {
		if err := f.Model(&favoriteModel).Where("video_id = ?", videoID).Update("favorite_count", count).Error; err != nil {
			return nil, err
		}
	}
	return counts, nil
}
