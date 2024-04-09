package main

import (
	"context"
	"github.com/oigi/Magikarp/app/comment/internal/model"
	"github.com/oigi/Magikarp/app/comment/internal/service"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/comment"
	"github.com/oigi/Magikarp/pkg/consts"
	"github.com/oigi/Magikarp/pkg/discovery"
	"github.com/oigi/Magikarp/pkg/loading"
	"github.com/oigi/Magikarp/pkg/mysql"
	"github.com/oigi/Magikarp/pkg/prometheus"
	"github.com/oigi/Magikarp/pkg/tracing"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	loading.Loading()
	mysql.InitMysql(&model.Comment{})

	// etcd 地址
	etcdAddress := []string{config.CONFIG.Etcd.Address}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress)

	grpcAddress := config.CONFIG.Etcd.Services[consts.CommentServiceName].Addr[0]

	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name: config.CONFIG.Etcd.Domain[consts.CommentServiceName].Name,
		Addr: grpcAddress,
	}
	//注册tracer
	provider := tracing.InitTracerProvider(config.CONFIG.Etcd.Jaeger.Addr, consts.CommentServiceName)
	defer func() {
		if provider == nil {
			return
		}
		if err := provider(context.Background()); err != nil {
			config.LOG.Error("Failed to shutdown:  ", zap.Error(err))
		}
	}()
	handler := otelgrpc.NewServerHandler()
	server := grpc.NewServer(
		grpc.StatsHandler(handler),
		grpc.UnaryInterceptor(prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(prometheus.StreamServerInterceptor),
	)
	defer server.Stop()
	// 绑定service
	comment.RegisterCommentServiceServer(server, service.GetCommentServe())
	prometheus.RegisterServer(server, config.CONFIG.Etcd.Services[consts.CommentServiceName].Metrics[0], consts.CommentServiceName)
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		config.LOG.Error("start service failed, original error:", zap.Error(err))
	}
	config.LOG.Info("service started listen on ", zap.String("address", grpcAddress))
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
