package rpc

import (
	"context"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
)

func ListFeed(ctx context.Context, req *feed.ListFeedReq) (resp *feed.ListFeedResp, err error) {
	stream, err := FeedClient.ListVideos(ctx, req)
	if err != nil {
		err = errors.WithMessage(err, "FeedClient.ListVideos error")
		config.LOG.Error("", zap.Error(err))
		return
	}
	var response *feed.ListFeedResp
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			err = errors.Wrap(err, "failed to receive response from ListVideos stream")
			config.LOG.Error("", zap.Error(err))
			return nil, err
		}
		response = resp
	}
	if response.Code != e.SUCCESS {
		err = errors.Wrap(errors.New("获取视频流失败"), "response code is unsuccessful")
		config.LOG.Error("", zap.Error(err))
		return nil, err
	}
	return response, nil

}

func QueryVideos(ctx context.Context, req *feed.QueryVideosReq) (resp *feed.QueryVideosResp, err error) {

	stream, err := FeedClient.QueryVideos(ctx)
	if err != nil {
		err = errors.WithMessage(err, "failed to initiate QueryVideos stream")
		config.LOG.Error("", zap.Error(err))
		return nil, err
	}

	err = stream.Send(req)
	if err != nil {
		err = errors.WithMessage(err, "failed to send QueryVideos request")
		config.LOG.Error("", zap.Error(err))
		return nil, err
	}

	var response *feed.QueryVideosResp

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			err = errors.Wrap(err, "failed to receive response from QueryVideos stream")
			config.LOG.Error("", zap.Error(err))
			return nil, err
		}
		response = resp
	}

	if err != nil {
		err = errors.WithMessage(err, "failed to receive response from QueryVideos stream")
		config.LOG.Error("", zap.Error(err))
		return nil, err
	}

	if response.Code != e.SUCCESS {
		err = errors.Wrap(errors.New("查询视频流失败"), "response code is unsuccessful")
		config.LOG.Error("", zap.Error(err))
		return nil, err
	}

	return response, nil
}
