package dao

import (
	"context"
	"github.com/oigi/Magikarp/app/feed/internal/consts"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
	"github.com/oigi/Magikarp/app/feed/internal/util"
	pkgRedis "github.com/oigi/Magikarp/pkg/redis"
	"github.com/pkg/errors"
	"time"
)

type RedisFeedDao struct {
	Client map[string]*pkgRedis.ClientRedis
	context.Context
}

func NewRedisClient(ctx context.Context) *RedisFeedDao {
	clients := pkgRedis.InitRedis()
	return &RedisFeedDao{
		clients,
		ctx,
	}
}

// GetFeedCache 从Redis中的用户feed信箱读取已经缓存过的视频列表
func (r *RedisFeedDao) GetFeedCache(ctx context.Context, email string, num int64) ([]feedModel.Videos, bool) {
	pops, err := r.Client[consts.FeedList].LPops(ctx, email, int(num))
	if err != nil {
		return nil, false
	}
	videoList, err := util.StringVideoList(pops)
	if err != nil {
		return nil, false
	}
	if len(videoList) > 10 {
		return videoList, true
	} else {
		return videoList, false
	}
}

// SetFeedCache 将视频列表缓存到Redis中的用户feed信箱
func (r *RedisFeedDao) SetFeedCache(ctx context.Context, method string, email string, videos ...feedModel.Videos) error {
	video_list := util.VideoListString(videos)
	switch method {
	case "l":
		return r.Client[consts.FeedList].LPush(ctx, email, video_list...)
	case "r":
		return r.Client[consts.FeedList].RPush(ctx, email, video_list...)
	default:
		return errors.New("unknown method, only accept 'l' or 'r'")
	}
}

// GetMarkedTime 从Redis中读取用户标记的时间
func (r *RedisFeedDao) GetMarkedTime(ctx context.Context, email string) (string, error) {
	return r.Client[consts.MongoMarkedTime].Get(ctx, email)
}

// SetMarkedTime 将用户标记的时间写入Redis
func (r *RedisFeedDao) SetMarkedTime(ctx context.Context, email, markedTime string) error {
	return r.Client[consts.MongoMarkedTime].Set(ctx, email, markedTime, 24*time.Hour)
}
