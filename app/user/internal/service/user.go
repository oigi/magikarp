package service

import (
    "context"
    "github.com/oigi/Magikarp/app/user/internal/dao"
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/consts/e"
    "github.com/oigi/Magikarp/grpc/pb/user"
    "github.com/oigi/Magikarp/initialize/mysql"
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
    client := dao.NewUserDao(ctx)
    defer mysql.CloseDB()
    r, err := client.GetUserInfo(req)
    if err != nil {
        resp.Code = e.ERROR
        config.LOG.Error("getUserInfo error", zap.Error(err))
        errors.WithMessage(err, "getUserInfo error")
        return
    }

    resp = &user.UserLoginResp{
        UserId: r.ID,
        Code:   e.SUCCESS,
        Msg:    "登陆成功",
        Email:  r.Email,
    }
    return
}
func (u *UserServe) UserRegister(ctx context.Context, req *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
    resp = new(user.UserRegisterResp)
    resp.Code = e.SUCCESS
    err = dao.NewUserDao(ctx).CreateUser(req)
    defer mysql.CloseDB()
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "注册失败"

    }
    return
}