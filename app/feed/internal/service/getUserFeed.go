package service

import (
	"context"
	"fmt"
	"github.com/oigi/Magikarp/app/feed/internal/dao"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
	"github.com/oigi/Magikarp/app/feed/internal/util"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"time"
)

// 1. 从 Redis 中获取用户 feed 列表
// 使用 LPop 命令从 Redis 列表中获取一定数量的 feed 数据
// 使用 Redis Pipeline 提高效率，减少网络开销
// 如果获取的数据不足以满足需求，则执行下一步骤

// 3. 处理视频数据
// 从 Redis 或 MySQL 中获取的视频数据进行处理
// 可以对视频数据进行排序、筛选或其他处理操作
// 选择一定数量的视频作为本次 feed 的数据，通常为前 10 条
// 将剩余的视频数据存储到 Redis 的投递箱中，用于后续的请求

// 4. 如果需要更新数据
// 计算当前时间与上次标记时间的差值，如果超过一定时间间隔，则执行查询操作
// 更新 Redis 中的用户 feed 数据，以确保数据的及时性和准确性

func (f *FeedServe) ListVideos(ctx context.Context, req *feed.ListFeedReq) (resp *feed.ListFeedResp, err error) {
	resp = &feed.ListFeedResp{}
	var logList []string

	mongoDao := dao.NewMongoClient(ctx)
	redis := dao.NewRedisClient(ctx)
	var list []feedModel.Videos
	var ok bool

	userIdString := util.FillUserId(fmt.Sprint(req.UserId))
	// 1. 从redis中获取用户feed列表 通过LPop
	req.LastTime = 0
	if req.LastTime == 0 {
		req.LastTime = time.Now().Unix()
	}

	if list, ok = redis.GetFeedCache(ctx, userIdString, 10); !ok {
		// 2. 【视频条数不足】从中从latest_time开始，以24h的周期向前查询，直至条数满足或超过current_time - 14 * 24h
		mongo, err := mongoDao.SearchFeedEarlierInMongo(req.LastTime, req.LastTime-14*24*60*60)
		if err != nil {
			return PackFeedListResp(ctx, []feedModel.Videos{}, 1, "search mongo failed", req.UserId)
		}
		if len(mongo) < 10 {
			req.LastTime = time.Now().Unix()
			mongo, err = mongoDao.SearchFeedEarlierInMongo(req.LastTime, req.LastTime-14*24*60*60)
			if err != nil {
				return PackFeedListResp(ctx, []feedModel.Videos{}, 1, "search mongo failed", req.UserId)
			}
		}

		// 3. 取前10条视频作为本次feed的数据，其余的通过RPush进入投递箱
		err = redis.SetFeedCache(ctx, "r", userIdString, mongo...)
		if err != nil {
			return PackFeedListResp(ctx, []feedModel.Videos{}, 1, "set send box failed", req.UserId)
		}
		var newListNum int64
		if len(mongo) > 10 {
			newListNum = 10
		} else {
			newListNum = int64(len(mongo))
		}

		list, ok = redis.GetFeedCache(ctx, userIdString, newListNum)
		//if !ok {
		//	return PackFeedListResp(ctx, []feedModel.Videos{}, 1, "get send box failed", req.UserId)
		//}
	}

	// 4. 计算current_time与marked_time的差值是否超过6个小时，如是则进行查询
	currentTime := time.Now().Unix()
	markedTime, err := redis.GetMarkedTime(ctx, userIdString)
	if err != nil {
		markedTime = fmt.Sprint(currentTime)
		err = redis.SetMarkedTime(ctx, userIdString, markedTime)
		if err != nil {
			logList = append(logList, "user_id为"+userIdString+"的用户设置marked_time失败")
		}
	}

	if util.JudgeTimeDiff(currentTime, markedTime, 60*60*6) {
		laterInMongo, newMarkedTime, err := mongoDao.SearchFeedLaterInMongo(markedTime, fmt.Sprint(currentTime))
		if err != nil {
			logList = append(logList, "user_id为"+userIdString+"的用户查询mongo失败")
		}

		err = redis.SetMarkedTime(ctx, userIdString, newMarkedTime)
		if err != nil {
			logList = append(logList, "user_id为"+userIdString+"的用户设置新的marked_time失败")
		}

		// 5. 若存在新更新的内容，将结果存入投递箱，根据比例选择RPush或LPush
		err = redis.SetFeedCache(ctx, "r", userIdString, laterInMongo...)
		if err != nil {
			logList = append(logList, "user_id为"+userIdString+"的用户设置send box失败")
		}
	}
	logString := ""
	for _, v := range logList {
		logString += v
		logString += ";  "
	}
	config.LOG.Warn(logString)

	return PackFeedListResp(ctx, list, 0, "Success", req.UserId)
}
