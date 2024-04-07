package mongo

import (
	"context"
	"github.com/oigi/Magikarp/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var mongoClient *mongo.Client

func InitMongoClient(ctx context.Context) *mongo.Client {
	if err := mgClient(ctx); err != nil {
		config.LOG.Error("Failed to initialize MongoDB client", zap.Error(err))
		return nil
	}
	return mongoClient
}

// mgClient 初始化 MongoDB 客户端
func mgClient(ctx context.Context) error {
	if mongoClient != nil {
		return nil // 如果已经初始化过了，则直接返回
	}

	uri := config.CONFIG.Mongo.Uri()

	if uri == "" {
		return errors.New("MongoDB URI is not set")
	}

	var err error
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	/*	coll := mongoClient.Database(m.Database).Collection(m.Collection)
		title := m.Title

		var result bson.M
		err = coll.FindOne(ctx, bson.D{{"title", title}}).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			config.LOG.Info("No document was found with the title", zap.String("title", title))
			return nil
		}
		if err != nil {
			return err
		}
		jsonData, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", jsonData)*/
	return nil
}

// CloseMongoClient 关闭 MongoDB 客户端连接
func CloseMongoClient(ctx context.Context) {
	if mongoClient != nil {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			config.LOG.Error("Failed to close MongoDB client", zap.Error(err))
		}
	}
}
