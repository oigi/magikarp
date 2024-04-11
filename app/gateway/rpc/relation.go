package rpc

import (
	"context"
	"github.com/oigi/Magikarp/grpc/pb/relation"
)

func RelationAction(ctx context.Context, req *relation.ActionReq) (resp *relation.ActionResp, err error) {
	return RelationClient.RelationAction(ctx, req)
}

func RelationFollowList(ctx context.Context, req *relation.FollowListReq) (resp *relation.FollowListResp, err error) {
	return RelationClient.RelationFollowList(ctx, req)
}

func RelationFollowerList(ctx context.Context, req *relation.FollowerListReq) (resp *relation.FollowerListResp, err error) {
	return RelationClient.RelationFollowerList(ctx, req)
}

func RelationFriendList(ctx context.Context, req *relation.FriendListReq) (resp *relation.FriendListResp, err error) {
	return RelationClient.RelationFriendList(ctx, req)
}
