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
	resp = &user.UserLoginResp{}
	client := dao.NewUserDao(ctx)
	//defer mysql.CloseDB() TODO 处理数据库关闭
	r, err := client.GetUserInfo(req)
	if err != nil {
		// 处理数据库查询错误
		config.LOG.Error("数据库查询错误", zap.Error(err))
		return
	}

	if r == nil {
		// 处理找不到用户信息的情况
		config.LOG.Error("找不到用户信息")
		resp.UserId = 0
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "找不到用户信息"
		return
	}

	// 处理正常情况
	userResp := &user.User{
		Id: r.ID,
	}
	resp.UserId = userResp.Id
	resp.StatusCode = e.SUCCESS
	resp.StatusMsg = "登录成功"
	return
}

func (u *UserServe) UserRegister(ctx context.Context, req *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
	resp = &user.UserRegisterResp{}
	id, err := dao.NewUserDao(ctx).CreateUser(req)
	if err != nil {
		config.LOG.Error("数据库查询错误", zap.Error(err))
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "注册失败"
		config.LOG.Error("createUser error", zap.Error(err))
		return resp, nil
	}
	resp.StatusCode = e.SUCCESS
	resp.StatusMsg = "注册成功"
	resp.UserId = id
	return
}

func (u *UserServe) GetUserById(ctx context.Context, req *user.GetUserByIdReq) (resp *user.GetUserByIdResp, err error) {
	r, err := dao.NewUserDao(ctx).GetUserInfoById(req)
	if err != nil {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "获取用户信息失败"
		return
	}
	userResp := queryDetailed(r)
	resp = &user.GetUserByIdResp{
		StatusCode: e.SUCCESS,
		StatusMsg:  "获取用户信息成功",
		User:       userResp,
	}
	return
}

func queryDetailed(r *userModel.User) (respUser *user.User) {
	respUser = &user.User{
		Id:            r.ID,
		Name:          r.NickName,
		FollowCount:   0, //TODO 需要修改
		FollowerCount: 0,
		IsFollow:      true, // TODO 需要修改
		//Sex:           int64(r.Sex),
		//Dec:           r.Dec,
		//Email:         r.Email,
		//Avatar:        r.Avatar,
	}
	return
}
