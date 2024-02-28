package main

import (
    "github.com/oigi/Magikarp/app/user/internal/service"
    "github.com/oigi/Magikarp/grpc/pb/user"
    "github.com/oigi/Magikarp/pkg/loading"
    "google.golang.org/grpc"
    "log"
    "net"
)

func main() {
    loading.Loading()

    grpcServer := grpc.NewServer()

    // 注册用户服务
    user.RegisterUserServiceServer(grpcServer, service.GetUserServe())

    // 监听网络端口
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 启动 gRPC 服务器
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
