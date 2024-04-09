package service

import (
    "context"
    "github.com/oigi/Magikarp/grpc/pb/relation"
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

}

func (r *RelationServe) RelationFollowList(ctx context.Context, req *relation.FollowListReq) (resp *relation.FollowListResp, err error) {

}

func (r *RelationServe) RelationFollowerList(ctx context.Context, req *relation.FollowerListReq) (resp *relation.FollowerListResp, err error) {

}

func (r *RelationServe) RelationFriendList(ctx context.Context, req *relation.FriendListReq) (resp *relation.FriendListResp, err error) {

}
