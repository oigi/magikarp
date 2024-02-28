package service

import (
    "context"
    "github.com/oigi/Magikarp/app/user/internal/dao"
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/consts/e"
    "github.com/oigi/Magikarp/grpc/pb/user"
    "github.com/oigi/Magikarp/pkg/jwt"
    "github.com/pkg/errors"
    "go.uber.org/zap"
    "sync"
)

var UserServeOnce sync.Once
var UserServeIns *UserServe

type UserServe struct {
    user.UnimplementedUserServiceServer
}

func GetUserServe() *UserServe {
    UserServeOnce.Do(func() {
        UserServeIns = &UserServe{}
    })
    return UserServeIns
}

func (u *UserServe) UserLogin(ctx context.Context, req *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
    resp = new(user.UserLoginResp)
    r, err := dao.NewUserDao(ctx).GetUserInfo(req)
    if err != nil {
        resp.Code = e.ERROR
        config.LOG.Error("getUserInfo error", zap.Error(err))
        errors.WithMessage(err, "getUserInfo error")
        return
    }
    // 生成 JWT token
    accessToken, refreshToken, err := jwt.GenerateJWT(213421, req.Email)
    if err != nil {
        config.LOG.Error("generate tokens error", zap.Error(err))
        return nil, errors.WithMessage(err, "generate tokens error")
    }

    resp = &user.UserLoginResp{
        Code:         e.SUCCESS,
        Msg:          "登陆成功",
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
        Email:        r.Email,
    }
    return
}
func (u *UserServe) UserRegister(ctx context.Context, req *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
    resp = new(user.UserRegisterResp)
    resp.Code = e.SUCCESS
    err = dao.NewUserDao(ctx).CreateUser(req)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "注册失败"

    }
    return
}
