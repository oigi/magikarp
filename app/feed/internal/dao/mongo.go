package dao

import (
	"context"
	"fmt"
	"github.com/oigi/Magikarp/app/feed/internal/consts"
	feedModel "github.com/oigi/Magikarp/app/feed/internal/model"
	"github.com/oigi/Magikarp/app/feed/internal/util"
	"github.com/oigi/Magikarp/config/model"
	mongodb "github.com/oigi/Magikarp/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
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

// GetVideoByUserIdInMongo 根据用户ID 在 MongoDB 中获取视频信息
func (m *MongoFeedDao) GetVideoByUserIdInMongo(UserID int) (*feedModel.Videos, error) {
	var video feedModel.Videos

	// 构建过滤条件
	filter := bson.D{
		{"user_id", UserID},
	}

	// 在 MongoDB 中搜索符合条件的视频
	err := m.Database(consts.MongoDatabaseName).Collection(consts.MongoCollection).FindOne(m.Context, filter).Decode(&video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

// FindFeedInMongo 在 MongoDB 中搜索 start_time < time <= end_time 的视频
func (m *MongoFeedDao) FindFeedInMongo(startTime int64, endTime int64) ([]feedModel.Videos, error) {
	// 构建过滤条件
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"timestamp", bson.D{{"$gt", startTime}}}},
				bson.D{{"timestamp", bson.D{{"$lte", endTime}}}},
			},
		},
	}

	// 在 MongoDB 中搜索符合条件的视频列表
	collection := m.Database(consts.MongoDatabaseName).Collection(consts.MongoCollection)
	cursor, err := collection.Find(m.Context, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// 遍历结果集，并将结果映射到 feedModel.Videos 结构体列表中
	var videos []feedModel.Videos
	for cursor.Next(context.Background()) {
		var video feedModel.VideoInMongo
		if err := cursor.Decode(&video); err != nil {
			continue
		}
		videomo := feedModel.Videos{
			Model: model.Model{
				ID: video.ID,
			},
			AuthorId:  video.UserID,
			Title:     video.Title,
			CoverUrl:  video.CoverURL,
			PlayUrl:   video.PlayURL,
			Category:  video.Category,
			Label:     video.Label,
			Timestamp: strconv.FormatInt(video.Timestamp, 10),
		}
		videos = append(videos, videomo)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

// SearchFeedEarlierInMongo 向前搜索 Feed List，前为更早的时间点
func (m *MongoFeedDao) SearchFeedEarlierInMongo(latestTime int64, stopTime int64) ([]feedModel.Videos, error) {
	var videoList []feedModel.Videos

	nextTime := latestTime - 86400

	for {
		videos, err := m.FindFeedInMongo(nextTime, latestTime)
		if err != nil {
			return videoList, err
		}

		videoList = append(videoList, videos...)

		// 终止条件1：视频列表长度已经大于30
		// 终止条件2：nextTime 小于 stopTime
		if len(videoList) > 30 || nextTime < stopTime {
			break
		}

		latestTime = nextTime
		nextTime -= 86400
	}

	return videoList, nil
}

// SearchFeedLaterInMongo 向后搜索 Feed List，后为更接近当前时间的时间点
func (m *MongoFeedDao) SearchFeedLaterInMongo(markedTime string, currentTime string) ([]feedModel.Videos, string, error) {
	markedTimeInt, _ := strconv.Atoi(markedTime)
	currentTimeInt, _ := strconv.Atoi(currentTime)

	nextMarkedTimeInt := int64(markedTimeInt) + 6*60*60

	var videoList []feedModel.Videos
	for {
		videos, err := m.FindFeedInMongo(int64(markedTimeInt), nextMarkedTimeInt)
		if err != nil {
			return videoList, markedTime, err
		}

		videoList = append(videoList, videos...)

		// 终止条件1：视频列表长度已经大于30
		// 终止条件2：时间差小于6个小时
		if len(videoList) > 30 || !util.JudgeTimeDiff(nextMarkedTimeInt, fmt.Sprint(currentTimeInt), 6*60*60) {
			break
		}

		markedTimeInt = int(nextMarkedTimeInt)
		nextMarkedTimeInt += 6 * 60 * 60
	}

	return videoList, fmt.Sprint(nextMarkedTimeInt), nil
}
