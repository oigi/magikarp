package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/user"
)

func GetUserById(ctx context.Context, req *user.GetUserByIdReq) (resp *user.GetUserByIdResp, err error) {
	return UserClient.GetUserById(ctx, req)
}
