package main

import (
    "context"
    "fmt"
    "github.com/oigi/Magikarp/app/gateway/routes"
    "github.com/oigi/Magikarp/app/gateway/rpc"
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/pkg/discovery"
    "github.com/oigi/Magikarp/pkg/loading"
    "github.com/oigi/Magikarp/pkg/tracing"
    "go.uber.org/zap"
    "google.golang.org/grpc/resolver"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    loading.Loading()
    rpc.Init()
    ////注册tracer
    provider := tracing.InitTracerProvider(config.CONFIG.Etcd.Jaeger.Addr, "magikarp")
    defer func() {
        if provider == nil {
            return
        }
        if err := provider(context.Background()); err != nil {
            config.LOG.Error("Failed to shutdown: ", zap.Error(err))
        }
    }()
    // etcd注册
    etcdAddress := []string{config.CONFIG.Etcd.Address}
    etcdRegister := discovery.NewResolver(etcdAddress)
    defer etcdRegister.Close()
    resolver.Register(etcdRegister)
    go startListen() // 转载路由
}

func startListen() {
    ginRouter := routes.NewRouter()
    server := &http.Server{
        Addr:           ":9090",
        Handler:        ginRouter,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    // 创建一个通道，用于接收关闭信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    // 启动 HTTP 服务器
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足 \n", ":9090")
            fmt.Println(err)
            return
        }
    }()

    fmt.Printf("gateway listen on :%v \n", "9090")

    // 等待接收关闭信号
    <-quit
    fmt.Println("Shutdown config ...")

    // 创建一个上下文，用于设置超时时间
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // 优雅关闭 HTTP 服务器
    if err := server.Shutdown(ctx); err != nil {
        fmt.Println("config Shutdown:", err)
    }

    fmt.Println("config exiting")
}
