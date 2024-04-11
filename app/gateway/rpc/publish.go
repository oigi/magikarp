package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/publish"
)

func CreateVideo(ctx context.Context, req *publish.CreateVideoRequest) (*publish.CreateVideoResponse, error) {
	return PublishClient.CreateVideo(ctx, req)
}

func ListVideo(ctx context.Context, req *publish.ListVideoRequest) (*publish.ListVideoResponse, error) {
	return PublishClient.ListVideo(ctx, req)
}
