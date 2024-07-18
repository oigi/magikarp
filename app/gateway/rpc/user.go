package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/user"
)

func UserLogin(ctx context.Context, req *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	return UserClient.UserLogin(ctx, req)
}

func UserRegister(ctx context.Context, req *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
	return UserClient.UserRegister(ctx, req)
}

func GetUserById(ctx context.Context, req *user.GetUserByIdReq) (resp *user.GetUserByIdResp, err error) {
	return UserClient.GetUserById(ctx, req)
}
