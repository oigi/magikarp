package service

import (
	"context"
	"github.com/oigi/Magikarp/app/comment/internal/dao"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/pkg/consts/e"
	"sync"
)

var CommentServeOnce sync.Once

var CommentServeIns *CommentServe

type CommentServe struct {
	comment.UnimplementedCommentServiceServer
}

func GetCommentServe() *CommentServe {
	CommentServeOnce.Do(func() {
		CommentServeIns = &CommentServe{}
	})
	return CommentServeIns
}

func (c *CommentServe) CommentAction(ctx context.Context, req *comment.CommentActionReq) (resp *comment.CommentActionResp, err error) {
	// 判断发布评论或者删除评论
	resp = &comment.CommentActionResp{}
	switch req.ActionType {
	case 1:
		if err = dao.NewCommentDao(ctx).AddComment(req); err != nil {
			resp.StatusCode = e.ERROR
			resp.StatusMsg = "评论失败"
		} else {
			resp.StatusCode = e.SUCCESS
			resp.StatusMsg = "评论成功"
		}
	case 2:
		if err = dao.NewCommentDao(ctx).DelComment(req); err != nil {
			resp.StatusCode = e.ERROR
			resp.StatusMsg = "删除评论失败"
		} else {
			resp.StatusCode = e.SUCCESS
			resp.StatusMsg = "删除评论成功"
		}
	}
	resp.StatusCode = e.InvalidParams
	resp.StatusMsg = "无效参数"
	return
}

func (c *CommentServe) CommentCount(ctx context.Context, req *comment.CommentCountReq) (resp *comment.CommentCountResp, err error) {
	resp = &comment.CommentCountResp{}
	count, err := dao.NewCommentDao(ctx).GetCommentCount(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = "获取评论数量失败"
		return
	}
	resp.CommentCount = count.CommentCount
	resp.Code = e.SUCCESS
	resp.Msg = "获取评论数量成功"
	return
}

func (c *CommentServe) CommentList(ctx context.Context, req *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	resp = &comment.CommentListResp{}
	list, err := dao.NewCommentDao(ctx).GetCommentList(req)
	if err != nil {
		resp.StatusCode = e.ERROR
		resp.StatusMsg = "获取评论列表失败"
		return
	}
	resp.CommentList = list.CommentList
	resp.StatusCode = e.SUCCESS
	resp.StatusMsg = "获取评论列表成功"
	return
}
