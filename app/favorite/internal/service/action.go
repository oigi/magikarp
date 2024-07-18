package service

import (
	"context"
	"github.com/oigi/Magikarp/app/favorite/internal/dao"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
)

// FavoriteAction 处理点赞操作 op 为 true 表示点赞，false 表示取消点赞
func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionReq, op bool) error {
	redis := dao.NewRedisClient(ctx)
	err := redis.WriteFavoriteInCache(ctx, req.UserId, req.VideoId, op)
	if err != nil {
		return err
	}
	err = redis.UpdateCacheCount(ctx, req.VideoId, op)
	if err != nil {
		return err
	}
	err = dao.NewMongoClient(context.Background()).WriteFavoriteInMongo(req.UserId, req.VideoId, op)
	if err != nil {
		return err
	}
	return nil
}
