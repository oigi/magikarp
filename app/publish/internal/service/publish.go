package service

import (
	"context"
	"github.com/oigi/Magikarp/app/publish/internal/dao"
	"github.com/oigi/Magikarp/app/publish/internal/util"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"go.uber.org/zap"
	"sync"
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

func (p *PublishServe) CreateVideo(ctx context.Context, req *publish.CreateVideoRequest) (*publish.CreateVideoResponse, error) {
	resp := &publish.CreateVideoResponse{}
	if ok, msg := util.Check(req); !ok {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = msg
		return resp, nil
	}
	err := SavePublish(ctx, req)
	if err != nil {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "创建视频失败"
		config.LOG.Error("创建视频失败", zap.Error(err))
		return resp, nil
	}
	resp.StatusCode = e.SUCCESS
	resp.StatusMsg = "创建视频成功"
	return resp, nil
}

func (p *PublishServe) ListVideo(ctx context.Context, req *publish.ListVideoRequest) (*publish.ListVideoResponse, error) {
	resp := &publish.ListVideoResponse{}

	list, err := dao.NewMongoClient(ctx).QueryPublishList(req.UserId)
	if err != nil {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "获取视频列表失败"
		return resp, err
	}

	resp.StatusCode = e.SUCCESS
	resp.StatusMsg = "获取视频列表成功"
	resp.VideoList = list
	return resp, nil
}

//func (p *PublishServe) DeleteVideo(ctx context.Context, req *publish.DeleteVideoReq) (*publish.DeleteVideoResp, error) {
//
//}
