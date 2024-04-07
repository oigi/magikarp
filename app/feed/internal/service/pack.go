package service

import (
	"context"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"go.uber.org/zap"
	"strconv"
)

func PackFeedListResp(vidoes []feedModel.Videos, code int64, msg string, userID int64, stream feed.Feed_ListVideosServer) error {
	nextTime := "7777777777"

	var VideoList []*feed.Video
	var VideoIdList []int64

	for _, v := range vidoes {
		VideoIdList = append(VideoIdList, v.ID)
	}

	isFavoriteResp, err := rpc.IsFavorite(context.Background(), &favorite.IsFavoriteReq{
		UserId:      userID,
		VideoIdList: VideoIdList,
	})
	if err != nil {
		config.LOG.Error("rpc.IsFavorite error", zap.Error(err))
		return nil
	}

	isFavorite := isFavoriteResp.IsFavorite

	favoriteCount, err := rpc.FavoriteCount(context.Background(), &favorite.FavoriteCountReq{
		VideoIdList: VideoIdList,
	})
	if err != nil {
		config.LOG.Error("rpc.FavoriteCount error", zap.Error(err))
		return nil
	}
	videoFavoriteCount := favoriteCount.VideoFavoriteCount

	commentCount, err := rpc.CommentCount(context.Background(), &comment.CommentCountReq{
		VideoId: VideoIdList,
	})
	if err != nil {
		config.LOG.Error("rpc.CommentCount error", zap.Error(err))
		return nil
	}
	videoCommentCount := commentCount.CommentCount

	for _, v := range vidoes {
		temp := &feed.Video{
			Id:            v.ID,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: videoFavoriteCount[v.ID],
			CommentCount:  videoCommentCount[v.ID],
			IsFavorite:    isFavorite[v.ID],
			Title:         v.Title,
			//StarCount:     0, TODO 加入收藏功能
			Duration: string(v.Duration),
			//PlayCount:    TODO 加入播放统计

		}
		resp, err := rpc.GetUserById(context.Background(), &user.GetUserByIdReq{
			UserId: v.AuthorId,
		})
		if err != nil {
			continue
		}

		temp.Author = resp.User
		VideoList = append(VideoList, temp)
		if v.Timestamp < nextTime {
			nextTime = v.Timestamp
		}
	}

	nextTimeInt, _ := strconv.Atoi(nextTime)
	stream.Send(&feed.ListFeedResp{
		Code:      code,
		Msg:       msg,
		VideoList: VideoList,
		NextTime:  int64(nextTimeInt),
	})

	return nil
}

func PackVideoInfoResp(v *feedModel.Videos) (*feed.Video, error) {
	video := &feed.Video{
		Id:       int64(v.ID),
		PlayUrl:  v.PlayUrl,
		CoverUrl: v.CoverUrl,
		Title:    v.Title,
		Duration: string(v.Duration),
	}

	idList := []int64{int64(v.ID)}
	favoriteCountResp, err := rpc.FavoriteCount(context.Background(), &favorite.FavoriteCountReq{
		VideoIdList: idList,
	})
	if err != nil {
		video.FavoriteCount = 0
	}
	commentCountResp, err := rpc.CommentCount(context.Background(), &comment.CommentCountReq{
		VideoId: idList,
	})
	if err != nil {
		video.CommentCount = 0
	}
	video.FavoriteCount = favoriteCountResp.VideoFavoriteCount[int64(v.ID)]
	video.CommentCount = commentCountResp.CommentCount[int64(v.ID)]
	video.IsFavorite = true

	return video, nil
}
