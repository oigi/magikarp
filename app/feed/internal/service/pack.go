package service

import (
	"context"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"go.uber.org/zap"
	"math"
)

func PackFeedListResp(vidoes []feedModel.Videos, code int64, msg string, userID int64, stream feed.Feed_ListVideosServer) error {
	resp := stream.Send(&feed.ListFeedResp{
		Code: code,
		Msg:  msg,
	})

	nextTime := math.MaxInt64

	var VideoidList []int64

	for _, v := range vidoes {
		VideoidList = append(VideoidList, v.ID)
	}

	isFavoriteResp, err := rpc.IsFavorite(context.Background(), &favorite.IsFavoriteReq{
		UserId:      userID,
		VideoIdList: VideoidList,
	})
	if err != nil {
		config.LOG.Error("rpc.IsFavorite error", zap.Error(err))
		return nil
	}

	isFavorite := isFavoriteResp.IsFavorite

	favoriteCount, err := rpc.FavoriteCount(context.Background(), &favorite.FavoriteCountReq{
		VideoIdList: VideoidList,
	})
	if err != nil {
		config.LOG.Error("rpc.FavoriteCount error", zap.Error(err))
		return nil
	}
	videoFavoriteCount := favoriteCount.VideoFavoriteCount

	commentCount, err := rpc.CommentCount(context.Background(), &comment.CommentCountReq{
		VideoId: VideoidList,
	})
	if err != nil {
		config.LOG.Error("rpc.CommentCount error", zap.Error(err))
		return nil
	}
	videoCommentCount := commentCount.CommentCount

	for _, v := range vidoes {
		temp := &feed.Video{
			Id:            v.ID,
			Uid:           v.AuthorId,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: videoFavoriteCount[v.ID],
			CommentCount:  videoCommentCount[v.ID],
			Title:         v.Title,
			CreateTime:    v.CreatedAt.Format("2006-01-02 15:04:05"),
			//StarCount:     0, TODO 加入收藏功能
			Duration: 	 string(v.Duration),
			//PlayCount:    TODO 加入播放统计
		}
	rpc.
	}

}
