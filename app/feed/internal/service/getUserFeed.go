package service

import (
	"fmt"
	"github.com/oigi/Magikarp/app/feed/internal/dao"
	"github.com/oigi/Magikarp/app/feed/internal/util"
	"github.com/oigi/Magikarp/grpc/pb/feed"
)

// 1. 从 Redis 中获取用户 feed 列表
// 使用 LPop 命令从 Redis 列表中获取一定数量的 feed 数据
// 使用 Redis Pipeline 提高效率，减少网络开销
// 如果获取的数据不足以满足需求，则执行下一步骤

// 2. 从 MySQL 中获取视频数据
// 从 MySQL 中根据条件查询视频数据，例如按发布时间倒序查询
// 可以根据用户 ID 和时间范围进行查询，并确保合适的索引被使用
// 如果获取的数据仍然不足以满足需求，则继续查询直至达到条件或超过时间范围

// 3. 处理视频数据
// 从 Redis 或 MySQL 中获取的视频数据进行处理
// 可以对视频数据进行排序、筛选或其他处理操作
// 选择一定数量的视频作为本次 feed 的数据，通常为前 10 条
// 将剩余的视频数据存储到 Redis 的投递箱中，用于后续的请求

// 4. 如果需要更新数据
// 计算当前时间与上次标记时间的差值，如果超过一定时间间隔，则执行查询操作
// 更新 Redis 中的用户 feed 数据，以确保数据的及时性和准确性

func (f *FeedServe) GetUserFeed(req *feed.ListFeedReq, stream feed.Feed_ListVideosServer) (
	err error) {
	userIdString := util.FillUserId(fmt.Sprint(req.ActorId))
	// 1. 从redis中获取用户feed列表 通过LPop

	mongoDao := dao.NewMongoClient(stream.Context())

	if list, ok := dao.GetFeedCache(stream.Context(), userIdString, 10); !ok {
		// 2. 【视频条数不足】从mysql中从latest_time开始，以24h的周期向前查询，直至条数满足或超过current_time - 14 * 24h
		mongo, err := mongoDao.SearchFeedEarlierInMongo(req.LastTime, req.LastTime-14*24*60*60)
		if err != nil {
			reture
		}
	}

	// 3. 取前10条视频作为本次feed的数据，其余的通过RPush进入投递箱

	// 4. 计算current_time与marked_time的差值是否超过6个小时，如是则进行查询

}
