package main

import (
    "fmt"
    "github.com/oigi/Magikarp/app/gateway/routes"
    "github.com/oigi/Magikarp/app/gateway/rpc"
    "github.com/oigi/Magikarp/pkg/loading"
    "net/http"
    "time"
)

func main() {
    loading.Loading()
    rpc.Init()
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
    if err := server.ListenAndServe(); err != nil {
        fmt.Printf("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足 \n", ":9090")
        fmt.Println(err)
        return
    }
    fmt.Printf("gateway listen on :%v \n", ":9090")
    // go func() {
    // 	// TODO 优雅关闭 有点问题，后续优化一下
    // 	shutdown.GracefullyShutdown(server)
    // }()
}
