package service

import (
	"github.com/oigi/Magikarp/app/feed/internal/dao"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"go.uber.org/zap"
	"sync"
	"time"
)

var FeedServeOnce sync.Once
var FeedServeIns *FeedServe

type FeedServe struct {
	feed.UnimplementedFeedServer
}

func GetFeedServe() *FeedServe {
	FeedServeOnce.Do(func() {
		FeedServeIns = &FeedServe{}
	})
	return FeedServeIns
}

func (f *FeedServe) ListVideos(req *feed.ListFeedReq, stream feed.Feed_ListVideosServer) (
	err error) {
	client := dao.NewFeedDao(stream.Context())
	find, err := client.FindVideos(req)
	if err != nil {
		config.LOG.Error("ListVideos video error", zap.Error(err))
		if err := stream.Send(&feed.ListFeedResp{
			Code: e.ERROR,
			Msg:  "ListVideos video error",
		}); err != nil {
			return err
		}
		return
	}

	var nextTime int64
	nextTime = find[len(find)-1].CreatedAt.Add(time.Duration(-1)).UnixMilli()

	videos := queryDetailed(find)

	if err := stream.Send(&feed.ListFeedResp{
		Code:      e.SUCCESS,
		Msg:       "find videos success",
		NextTime:  nextTime,
		VideoList: videos,
	}); err != nil {
		return err
	}
	return
}

func (f *FeedServe) QueryVideos(req *feed.QueryVideosReq, stream feed.Feed_QueryVideosServer) (
	err error) {
	client := dao.NewFeedDao(stream.Context())
	find, err := client.FindVideosByUser(req)
	if err != nil {
		return err
	}
	videos := queryDetailed(find)

	// 将视频数据发送给客户端流
	if err := stream.Send(&feed.QueryVideosResp{
		Code:      200,
		Msg:       "传输成功",
		VideoList: videos,
	}); err != nil {
		return err
	}

	return
}

func queryDetailed(videos []feedModel.Videos) (respVideoList []*feed.Video) {
	var wg sync.WaitGroup
	videoChan := make(chan *feed.Video, len(videos))

	for _, v := range videos {
		wg.Add(1)
		go func(v feedModel.Videos) {
			defer wg.Done()
			videoChan <- &feed.Video{
				Id:            v.ID,
				Uid:           v.AuthorId,
				PlayUrl:       v.PlayUrl,
				CoverUrl:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount:  v.CommentCount,
				Title:         v.Title,
				CreateTime:    v.CreatedAt.Format(time.RFC3339), // 格式化时间
				StarCount:     v.StarCount,
				Duration:      v.Duration.String(), // Duration转换为字符串
				PlayCount:     v.PlayCount,
			}
		}(v)
	}
	wg.Wait()
	close(videoChan)

	for video := range videoChan {
		respVideoList = append(respVideoList, video)
	}

	return respVideoList
}
