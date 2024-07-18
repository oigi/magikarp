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

func InitMongoClient() *mongo.Client {
	if err := mgClient(); err != nil {
		config.LOG.Error("Failed to initialize MongoDB client", zap.Error(err))
		return nil
	}
	return mongoClient
}

// mgClient 初始化 MongoDB 客户端
func mgClient() error {
	if mongoClient != nil {
		return nil // 如果已经初始化过了，则直接返回
	}

	uri := config.CONFIG.Mongo.Uri()

	if uri == "" {
		return errors.New("MongoDB URI is not set")
	}

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	return nil
}

// CloseMongoClient 关闭 MongoDB 客户端连接
func CloseMongoClient() {
	if mongoClient != nil {
		err := mongoClient.Disconnect(context.Background())
		if err != nil {
			config.LOG.Error("Failed to close MongoDB client", zap.Error(err))
		}
	}
}
