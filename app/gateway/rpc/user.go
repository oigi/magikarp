package rpc

import (
	"context"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func UserLogin(ctx context.Context, req *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	r, err := UserClient.UserLogin(ctx, req)
	if err != nil {
		err = errors.WithMessage(err, "UserClient.UserLogin error")
		config.LOG.Error("", zap.Error(err))
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.Wrap(errors.New("登陆失败"), "r.Code is unsuccessful")
		config.LOG.Error("", zap.Error(err))
		return
	}

	return r, nil
}

func UserRegister(ctx context.Context, req *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
	r, err := UserClient.UserRegister(ctx, req)
	if err != nil {
		err = errors.WithMessage(err, "UserClient.UserRegister error")
		config.LOG.Error("", zap.Error(err))
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.Wrap(errors.New(r.Msg), "r.Code is unsuccessful")
		config.LOG.Error("", zap.Error(err))
		return
	}

	return
}

func GetUserById(ctx context.Context, req *user.GetUserByIdReq) (resp *user.GetUserByIdResp, err error) {
	r, err := UserClient.GetUserById(ctx, req)
	if err != nil {
		err = errors.WithMessage(err, "UserClient.GetUserById error")
		config.LOG.Error("", zap.Error(err))
		return
	}

	if r.Code != e.SUCCESS {
		err = errors.Wrap(errors.New(r.Msg), "r.Code is unsuccessful")
		config.LOG.Error("", zap.Error(err))
		return
	}

	return
}
