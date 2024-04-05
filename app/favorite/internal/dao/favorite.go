package dao

import (
	"context"
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

//func (f *FavoriteDao) UpdateVideos(req *favorite.SyncFavoriteReq) (err error)  {
//
//}
