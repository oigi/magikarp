package prometheus

import (
	"encoding/json"
	"fmt"
	"github.com/oigi/Magikarp/config"
	"go.uber.org/zap"
	"os"
)

// GenerateAllConfigFile generate configuration files
// for all registered services
func GenerateAllConfigFile() {
	service := config.CONFIG.Etcd.Services
	if len(service) == 0 {
		return
	}
	for k, _ := range service {
		GenerateConfigFile(k)
	}
}

// GenerateConfigFile generate configuration files
// for the services
func GenerateConfigFile(job string) {
	instance := GetServerAddress(job)

	f, err := os.OpenFile(fmt.Sprintf("./pkg/prometheus/model/files/%s.json", job), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		config.LOG.Error(fmt.Sprintf("failed open file prometheus/model/files/%s.json", job), zap.Error(err))
		return
	}
	defer f.Close()
	buf, err := json.MarshalIndent(instance.Conf, "", "    ")
	if err != nil {
		config.LOG.Error("failed marshal", zap.Error(err))
		return
	}
	_, err = f.Write(buf)
	if err != nil {
		config.LOG.Error("failed write to file", zap.Error(err))
		return
	}
}
