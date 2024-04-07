package service

import (
	"context"
	"github.com/oigi/Magikarp/app/publish/internal/dao"
	publishModel "github.com/oigi/Magikarp/app/publish/internal/model"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"sync"
	"time"
)

var PublishServeOnce sync.Once
var PublishServeIns *PublishServe

type PublishServe struct {
	publish.UnimplementedPublishServiceServer
}

func GetPublishServe() *PublishServe {
	PublishServeOnce.Do(func() {
		PublishServeIns = &PublishServe{}
	})
	return PublishServeIns
}

func (p *PublishServe) CreateVideo(ctx context.Context, req *publish.CreateVideoRequest) (resp *publish.CreateVideoResponse, err error) {
	resp.Code = e.SUCCESS
	client := dao.NewPublishDao(ctx)
	err = client.InsertVideo(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = "创建视频失败"
		return
	}
	resp.Code = 200
	resp.Msg = "创建视频成功"
	return
}
func (p *PublishServe) ListVideo(ctx context.Context, req *publish.ListVideoRequest) (resp *publish.ListVideoResponse, err error) {
	find, err := dao.NewPublishDao(ctx).FindVideoListByUserId(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = "查询用户视频失败"
		return
	}
	videos := queryDetailed(find)
	resp.VideoList = videos
	resp.Code = e.SUCCESS
	resp.Msg = "查询视频成功"
	return
}

//func (p *PublishServe) CountVideo(ctx context.Context, req *publish.CountVideoRequest) (resp *publish.CountVideoResponse, err error) {
//
//}

func queryDetailed(videos []publishModel.Videos) (respVideoList []*feed.Video) {
	var wg sync.WaitGroup
	videoChan := make(chan *feed.Video, len(videos))

	for _, v := range videos {
		wg.Add(1)
		go func(v publishModel.Videos) {
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
