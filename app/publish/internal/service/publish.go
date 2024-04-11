package service

import (
	"context"
	"github.com/oigi/Magikarp/app/publish/internal/dao"
	"github.com/oigi/Magikarp/app/publish/internal/util"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"github.com/oigi/Magikarp/pkg/consts/e"
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
		resp.Code = e.ERROR
		resp.Msg = msg
		return resp, nil
	}
	err := SavePublish(ctx, req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = "创建视频失败"
		return resp, err
	}
	resp.Code = e.SUCCESS
	resp.Msg = "创建视频成功"
	return resp, nil
}

func (p *PublishServe) ListVideo(ctx context.Context, req *publish.ListVideoRequest) (*publish.ListVideoResponse, error) {
	resp := &publish.ListVideoResponse{}

	list, err := dao.NewMongoClient(ctx).QueryPublishList(req.UserId)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = "获取视频列表失败"
		return resp, err
	}

	resp.Code = e.SUCCESS
	resp.Msg = "获取视频列表成功"
	resp.VideoList = list
	return resp, nil
}

//func (p *PublishServe) DeleteVideo(ctx context.Context, req *publish.DeleteVideoReq) (*publish.DeleteVideoResp, error) {
//
//}
