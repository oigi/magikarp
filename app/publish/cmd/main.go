package main

import (
	"context"
	"github.com/oigi/Magikarp/app/publish/internal/model"
	"github.com/oigi/Magikarp/app/publish/internal/rpc"
	"github.com/oigi/Magikarp/app/publish/internal/service"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/grpc/pb/publish"
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
	mysql.InitMysql(&model.Video{})
	rpc.Init()
	// etcd 地址
	etcdAddress := []string{config.CONFIG.Etcd.Address}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress)

	grpcAddress := config.CONFIG.Etcd.Services[consts.PublishServiceName].Addr[0]

	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name: config.CONFIG.Etcd.Domain[consts.PublishServiceName].Name,
		Addr: grpcAddress,
	}
	//注册tracer
	provider := tracing.InitTracerProvider(config.CONFIG.Etcd.Jaeger.Addr, consts.PublishServiceName)
	defer func() {
		if provider == nil {
			return
		}
		if err := provider(context.Background()); err != nil {
			config.LOG.Error("Failed to shutdown:  ", zap.Error(err))
		}
	}()
	handler := otelgrpc.NewServerHandler()
	mb := 1024 * 1024
	server := grpc.NewServer(
		grpc.StatsHandler(handler),
		grpc.UnaryInterceptor(prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(prometheus.StreamServerInterceptor),
		grpc.MaxRecvMsgSize(100*mb),
		grpc.MaxSendMsgSize(100*mb),
	)
	defer server.Stop()
	// 绑定service
	publish.RegisterPublishServiceServer(server, service.GetPublishServe())
	prometheus.RegisterServer(server, config.CONFIG.Etcd.Services[consts.PublishServiceName].Metrics[0], consts.PublishServiceName)
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
