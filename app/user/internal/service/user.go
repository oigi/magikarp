package service

import (
	"context"
	"github.com/oigi/Magikarp/app/user/internal/dao"
	userModel "github.com/oigi/Magikarp/app/user/internal/model"
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
	userResp := queryDetailed(r)
	resp = &user.UserLoginResp{
		User: userResp,
		Code: e.SUCCESS,
		Msg:  "登陆成功",
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

func queryDetailed(r *userModel.User) (respUser *user.User) {
	respUser = &user.User{
		Id:            r.ID,
		Name:          r.NickName,
		FollowCount:   int64(r.FollowCount),
		FollowerCount: int64(r.FollowerCount),
		Sex:           int64(r.Sex),
		Dec:           r.Dec,
		Email:         r.Email,
		Avatar:        r.Avatar,
	}
	return
}
