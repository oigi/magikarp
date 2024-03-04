package prometheus

import (
    "github.com/gin-gonic/gin"
    "github.com/oigi/Magikarp/config"
    "github.com/oigi/Magikarp/pkg/consts"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "go.uber.org/zap"
    "net/http"
    "strings"
)

// GatewayHandler for expose gateway metrics
func GatewayHandler() gin.HandlerFunc {
    EtcdRegister(config.CONFIG.Etcd.Server.Metrics, consts.GatewayJobForPrometheus)
    handler := promhttp.Handler()
    return func(c *gin.Context) {
        handler.ServeHTTP(c.Writer, c.Request)
    }
}

// RpcHandler is for launch a http server
// to expose metrics
func RpcHandler(addr string) {
    port := strings.Split(addr, ":")[1]
    http.Handle("/metrics", promhttp.Handler())
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        config.LOG.Panic("Failed to start server", zap.Error(err))
    }

}
