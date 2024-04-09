package service

import (
    "context"
    "github.com/oigi/Magikarp/app/feed/internal/dao"
    "github.com/oigi/Magikarp/grpc/pb/feed"
)

func (f *FeedServe) GetVideoById(ctx context.Context, req *feed.QueryVideosReq) (resp *feed.Video, err error) {
    mongoDao := dao.NewMongoClient(ctx)
    date, err := mongoDao.GetVideoByUserIdInMongo(int(req.VideoId))
    if err != nil {
        return &feed.Video{}, nil
    }
    return PackVideoInfoResp(date)
}
