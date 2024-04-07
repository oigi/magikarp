package dao

import (
	"context"
	"github.com/oigi/Magikarp/app/feed/internal/consts"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
	"github.com/oigi/Magikarp/app/feed/internal/util"
	"github.com/oigi/Magikarp/pkg/redis"
	"github.com/pkg/errors"
	"time"
)

var RedisClient map[string]*redis.ClientRedis

// GetFeedCache 从Redis中的用户feed信箱读取已经缓存过的视频列表
func GetFeedCache(ctx context.Context, email string, num int64) ([]feedModel.Videos, bool) {
	pops, err := RedisClient[consts.FeedList].LPops(ctx, email, int(num))
	if err != nil {
		return nil, false
	}

	videoList, err := util.StringVideoList(pops)
	if err != nil {
		return nil, false
	}

	return videoList, true
}

// SetFeedCache 将视频列表缓存到Redis中的用户feed信箱
func SetFeedCache(ctx context.Context, method string, email string, videos ...feedModel.Videos) error {
	video_list := util.VideoListString(videos)
	switch method {
	case "l":
		return RedisClient[consts.FeedList].LPush(ctx, email, video_list...)
	case "r":
		return RedisClient[consts.FeedList].RPush(ctx, email, video_list...)
	default:
		return errors.New("unknown method, only accept 'l' or 'r'")
	}
}

// GetMarkedTime 从Redis中读取用户标记的时间
func GetMarkedTime(ctx context.Context, email string) (string, error) {
	return RedisClient[consts.MongoMarkedTime].Get(ctx, email)
}

// SetMarkedTime 将用户标记的时间写入Redis
func SetMarkedTime(ctx context.Context, email, markedTime string) error {
	return RedisClient[consts.MongoMarkedTime].Set(ctx, email, markedTime, 24*time.Hour)
}
