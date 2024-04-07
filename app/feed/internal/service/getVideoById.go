package service

import (
	"github.com/oigi/Magikarp/app/feed/internal/dao"
	"github.com/oigi/Magikarp/grpc/pb/feed"
)

func (f *FeedServe) QueryVideos(req *feed.QueryVideosReq, stream feed.Feed_QueryVideosServer) (
	err error) {
	mongoDao := dao.NewMongoClient(stream.Context())
	date, err := mongoDao.GetVideoByUserIdInMongo(int(req.ActorId))
	if err != nil {
		return err
	}
	resp, _ := PackVideoInfoResp(date)
	if err := stream.SendMsg(resp); err != nil {
		return err
	}
	return nil
}
