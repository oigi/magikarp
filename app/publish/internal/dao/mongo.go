package dao

import (
	"context"
	"fmt"
	"github.com/oigi/Magikarp/app/gateway/rpc"
	"github.com/oigi/Magikarp/app/publish/internal/consts"
	publishModel "github.com/oigi/Magikarp/app/publish/internal/model"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/grpc/pb/user"
	mongodb "github.com/oigi/Magikarp/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFeedDao struct {
	*mongo.Client
	context.Context
}

func NewMongoClient(ctx context.Context) *MongoFeedDao {
	return &MongoFeedDao{
		mongodb.InitMongoClient(),
		ctx,
	}
}

func (m *MongoFeedDao) InsertVideoInMongo(
	id int64, userId int64, title string, playUrl string, coverUrl string, label string, category string, timestamp string) error {
	video := publishModel.VideoInMongo{
		ID:        id,
		UserID:    userId,
		Title:     title,
		PlayURL:   playUrl,
		CoverURL:  coverUrl,
		Label:     label,
		Category:  category,
		Timestamp: timestamp,
	}
	collection := m.Database(consts.MongoDatabaseName).Collection(consts.MongoCollection)

	_, err := collection.InsertOne(m.Context, video)
	if err != nil {
		return fmt.Errorf("failed to insert video: %v", err)
	}

	return nil
}

func (m *MongoFeedDao) QueryPublishList(userId int64) ([]*feed.Video, error) {
	var videosInMongo []publishModel.VideoInMongo
	var videosFeed []*feed.Video
	filter := bson.M{"userId": userId}

	collection := m.Database(consts.MongoDatabaseName).Collection(consts.MongoCollection)

	cur, err := collection.Find(m.Context, filter)
	if err != nil {
		return nil, fmt.Errorf("查询视频失败: %v", err)
	}
	defer cur.Close(m.Context)

	for cur.Next(m.Context) {
		var video publishModel.VideoInMongo
		err := cur.Decode(&video)
		if err != nil {
			return nil, fmt.Errorf("解码视频失败: %v", err)
		}
		videosInMongo = append(videosInMongo, video)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("游标错误: %v", err)
	}

	// 调用rpc.GetUserInfoListByIds获取用户信息
	req := user.GetUserByIdReq{UserId: userId}
	user, err := rpc.GetUserById(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	// 遍历videosInMongo，将其转换为feed.Video格式
	for _, v := range videosInMongo {
		videoFeed := &feed.Video{
			Id:       v.ID, // 假设v.ID可以直接赋值
			PlayUrl:  v.PlayURL,
			CoverUrl: v.CoverURL,
			Author:   user.User,
			//FavoriteCount: 0,
			//CommentCount:  0,
			Title: v.Title,
			//StarCount:     0,
			//Duration:      "",
			//PlayCount:     0,
			//IsFavorite:    false,
		}
		videosFeed = append(videosFeed, videoFeed)
	}

	return videosFeed, nil
}
