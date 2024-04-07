package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/pkg/errors"
)

func CommentCount(ctx context.Context, req *comment.CommentCountReq) (resp *comment.CommentCountResp, err error) {
	resp, err = CommentClient.CommentCount(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	return resp, nil
}
