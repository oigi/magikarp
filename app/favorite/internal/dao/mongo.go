package dao

import (
    "context"
    "github.com/oigi/Magikarp/app/favorite/internal/consts"
    mongodb "github.com/oigi/Magikarp/pkg/mongo"
    "github.com/pkg/errors"
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

func (m *MongoFeedDao) WriteFavoriteInMongo(userId int64, videoId int64, isFavorite bool) error {
    collection := m.Database(consts.MongoDatabaseName).Collection(consts.MongoCollection)
    filter := bson.M{"user_id": userId, "video_id": videoId}

    var existing bson.M
    err := collection.FindOne(m.Context, filter).Decode(&existing)
    if err == nil {
        // 如果记录已存在，则更新favorite字段
        existing["favorite"] = isFavorite
        _, err := collection.ReplaceOne(m.Context, filter, existing)
        return err
    } else if !errors.Is(err, mongo.ErrNoDocuments) {
        return err
    }

    doc := bson.M{
        "user_id":  userId,
        "video_id": videoId,
        "favorite": isFavorite,
    }

    _, err = collection.InsertOne(m.Context, doc)
    return err
}

func (m *MongoFeedDao) QueryFavoriteListInMongo(userId int64) ([]int64, error) {
    collection := m.Database(consts.MongoDatabaseName).Collection(consts.MongoCollection)
    filter := bson.M{"user_id": userId}

    cursor, err := collection.Find(m.Context, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(m.Context)

    var results []int64
    for cursor.Next(m.Context) {
        var result bson.M
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
        videoId, ok := result["video_id"].(int64)
        if !ok {
            return nil, errors.New("video_id is not an int64")
        }
        results = append(results, videoId)
    }
    return results, nil
}
