package service

import (
    "context"
    "github.com/oigi/Magikarp/app/user/internal/dao"
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/grpc/pb/user"
    "github.com/oigi/Magikarp/pkg/consts/e"
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
    client := dao.NewUserDao(ctx)
    //defer mysql.CloseDB() TODO 处理数据库关闭
    r, err := client.GetUserInfo(req)
    if err != nil {
        resp.Code = e.ERROR
        config.LOG.Error("getUserInfo error", zap.Error(err))
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
    resp.Code = e.SUCCESS
    err = dao.NewUserDao(ctx).CreateUser(req)
    if err != nil {
        resp.Code = e.ERROR
        resp.Msg = "注册失败"
    }
    return
}
