package service

import (
	"context"
	"github.com/oigi/Magikarp/app/relation/internal/dao"
	"github.com/oigi/Magikarp/grpc/pb/relation"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"sync"
)

var RelationServeOnce sync.Once
var RelationServeIns *RelationServe

type RelationServe struct {
	relation.UnimplementedRelationServiceServer
}

func GetRelationServe() *RelationServe {
	RelationServeOnce.Do(func() {
		RelationServeIns = &RelationServe{}
	})
	return RelationServeIns
}

func (r *RelationServe) RelationAction(ctx context.Context, req *relation.ActionReq) (resp *relation.ActionResp, err error) {
	resp = &relation.ActionResp{}
	switch req.ActionType {
	case 1:
		if err = dao.NewRelationDao(ctx).AddRelation(req); err != nil {
			resp.StatusCode = e.ERROR
			resp.StatusMsg = "添加关注失败"
			return
		} else {
			resp.StatusCode = e.SUCCESS
			resp.StatusMsg = "添加关注成功"
			return
		}
	case 2:
		if err = dao.NewRelationDao(ctx).DeleteRelation(req); err != nil {
			resp.StatusCode = e.ERROR
			resp.StatusMsg = "取消关注失败"
			return
		} else {
			resp.StatusCode = e.SUCCESS
			resp.StatusMsg = "取消关注成功"
			return
		}
	}
	resp.StatusCode = e.InvalidParams
	resp.StatusMsg = "无效的操作类型"
	return
}

func (r *RelationServe) RelationFollowList(ctx context.Context, req *relation.FollowListReq) (resp *relation.FollowListResp, err error) {
	resp = &relation.FollowListResp{}
	if list, err := dao.NewRelationDao(ctx).QueryFollowingUsers(req.UserId); err != nil {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "获取关注列表失败"
	} else {
		resp.StatusCode = e.SUCCESS
		resp.StatusMsg = "获取关注列表成功"
		resp.FollowList = list
	}
	return
}

func (r *RelationServe) RelationFollowerList(ctx context.Context, req *relation.FollowerListReq) (resp *relation.FollowerListResp, err error) {
	resp = &relation.FollowerListResp{}
	if list, err := dao.NewRelationDao(ctx).QueryFollowers(req.UserId); err != nil {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "获取粉丝列表失败"
	} else {
		resp.StatusCode = e.SUCCESS
		resp.StatusMsg = "获取粉丝列表成功"
		resp.FollowerList = list
	}
	return
}

func (r *RelationServe) RelationFriendList(ctx context.Context, req *relation.FriendListReq) (resp *relation.FriendListResp, err error) {
	resp = &relation.FriendListResp{}
	if list, err := dao.NewRelationDao(ctx).QueryMutualFriends(req.UserId); err != nil {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "获取好友列表失败"
	} else {
		resp.StatusCode = e.SUCCESS
		resp.StatusMsg = "获取好友列表成功"
		resp.FriendList = list
	}
	return
}
