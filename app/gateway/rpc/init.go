package rpc

import (
	"context"
	"fmt"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/grpc/pb/favorite"
	"github.com/oigi/Magikarp/grpc/pb/feed"
	"github.com/oigi/Magikarp/grpc/pb/publish"
	"github.com/oigi/Magikarp/grpc/pb/relation"
	"github.com/oigi/Magikarp/grpc/pb/user"
	"github.com/oigi/Magikarp/pkg/consts"
	"github.com/oigi/Magikarp/pkg/discovery"
	"github.com/oigi/Magikarp/pkg/prometheus"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var (
	Register   *discovery.Resolver
	ctx        context.Context
	CancelFunc context.CancelFunc

	UserClient     user.UserServiceClient
	FeedClient     feed.FeedClient
	FavoriteClient favorite.FavoriteServiceClient
	CommentClient  comment.CommentServiceClient
	PublishClient  publish.PublishServiceClient
	RelationClient relation.RelationServiceClient
)

// Init 初始化所有的rpc请求
func Init() {
	Register = discovery.NewResolver([]string{config.CONFIG.Etcd.Address})
	resolver.Register(Register)
	//ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()
	initClient(config.CONFIG.Etcd.Domain[consts.UserServiceName].Name, &UserClient)
	initClient(config.CONFIG.Etcd.Domain[consts.FeedServiceName].Name, &FeedClient)
	initClient(config.CONFIG.Etcd.Domain[consts.FavoriteServiceName].Name, &FavoriteClient)
	initClient(config.CONFIG.Etcd.Domain[consts.CommentServiceName].Name, &CommentClient)
	initClient(config.CONFIG.Etcd.Domain[consts.PublishServiceName].Name, &PublishClient)
	initClient(config.CONFIG.Etcd.Domain[consts.RelationServiceName].Name, &RelationClient)
}

// initClient 初始化所有的rpc客户端
func initClient(serviceName string, client interface{}) {
	prometheus.EnableHandlingTimeHistogram()
	conn, err := connectServer(serviceName)

	if err != nil {
		config.LOG.Panic("", zap.Error(err))
		panic(err)
	}

	switch c := client.(type) {
	case *user.UserServiceClient:
		*c = user.NewUserServiceClient(conn)
	case *favorite.FavoriteServiceClient:
		*c = favorite.NewFavoriteServiceClient(conn)
	case *feed.FeedClient:
		*c = feed.NewFeedClient(conn)
	case *comment.CommentServiceClient:
		*c = comment.NewCommentServiceClient(conn)
	case *publish.PublishServiceClient:
		*c = publish.NewPublishServiceClient(conn)
	case *relation.RelationServiceClient:
		*c = relation.NewRelationServiceClient(conn)

	default:
		panic("unsupported worker type")
	}
}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)

	// Load balance
	if config.CONFIG.Etcd.Services[serviceName].LoadBalance {
		log.Printf("load balance enabled for %s\n", serviceName)
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))
	}

	conn, err = grpc.DialContext(ctx, addr, opts...)
	err = errors.Wrapf(err, "failed to connect to gRPC service,address is %v", addr)
	return
}
