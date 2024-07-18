package dao

import (
	"context"
	"fmt"
	"github.com/oigi/Magikarp/app/favorite/internal/consts"
	"github.com/oigi/Magikarp/pkg/redis"
	"strconv"
)

type RedisFavoriteDao struct {
	Client map[string]*redis.ClientRedis
	context.Context
}

func NewRedisClient(ctx context.Context) *RedisFavoriteDao {
	clients := redis.InitRedis()
	return &RedisFavoriteDao{
		clients,
		ctx,
	}
}

func (r *RedisFavoriteDao) WriteFavoriteInCache(ctx context.Context, userId int64, videoId int64, isFavorite bool) error {
	rdb := r.Client[consts.FavCache]
	var status string
	if isFavorite {
		status = "1" // 已点赞
	} else {
		status = "0" // 未点赞
	}
	return rdb.HSet(ctx, fmt.Sprint(userId), fmt.Sprint(videoId), status)
}

func (r *RedisFavoriteDao) UpdateCacheCount(ctx context.Context, videoId int64, isFavorite bool) error {
	rdb := r.Client[consts.FavCache]
	var status int
	if isFavorite {
		status = 1
	} else {
		status = -1
	}
	result, err := rdb.Get(ctx, fmt.Sprint(videoId))
	if err != nil {
		return rdb.Client.Set(ctx, fmt.Sprint(videoId), "1", -1).Err()
	}
	atoi, _ := strconv.Atoi(result)
	atoi += status

	return rdb.Client.Set(context.Background(), fmt.Sprint(videoId), fmt.Sprint(atoi), -1).Err()
}

// QueryFavoriteListInCache 查询用户点赞列表
func (r *RedisFavoriteDao) QueryFavoriteListInCache(ctx context.Context, userId int64) ([]int64, error) {
	rdb := r.Client[consts.FavCache]
	m, err := rdb.Get(ctx, fmt.Sprint(userId))
	if err != nil {
		return nil, err
	}
	var result []int64
	for k, v := range m {
		kInt64 := int64(k)
		videoId := kInt64
		if err != nil {
			return nil, err
		}
		if v == 1 {
			result = append(result, videoId)
		}
	}
	return result, nil
}

func (r *RedisFavoriteDao) QueryFavoriteCount(videoIds []int64) (map[int64]int64, error) {
	rdb := r.Client[consts.FavCache]
	counts := make(map[int64]int64)

	// 查询点赞数量
	for _, videoId := range videoIds {
		countStr, _ := rdb.Get(context.Background(), strconv.FormatInt(videoId, 10))
		if countStr != "" {
			count, err := strconv.ParseInt(countStr, 10, 64)
			if err != nil {
				return nil, err
			}
			counts[videoId] = count
			continue
		}
		counts[videoId] = 0
	}

	return counts, nil
}
