package prometheus

import (
	"context"
	"fmt"
	"github.com/oigi/Magikarp/config"
	"github.com/oigi/Magikarp/consts"
	"go.uber.org/zap"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

// Instance is for marshal conf
type Instance struct {
	Conf []*Conf
}

// Conf is the basic unit of the prometheus detection unit
type Conf struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}

// EtcdRegister need server address and name
// for register to etcd and keep alive
func EtcdRegister(targets string, job string) {
	client := newClient()
	leaseResp, err := client.Grant(context.Background(), 15)
	if err != nil {
		config.LOG.Error("", zap.Error(err))
	}
	key := fmt.Sprintf("%s/%s/%d", consts.PrometheusJobKey, job, leaseResp.ID)
	if _, err = client.Put(context.Background(), key, targets, etcd.WithLease(leaseResp.ID)); err != nil {
		config.LOG.Error("", zap.Error(err))
		return
	}

	go keepALive(client, leaseResp.ID)
	go GenerateConfigFile(job)
}

// keepAlive for registered instance
func keepALive(c *etcd.Client, leaseId etcd.LeaseID) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	keepLiveCh, _ := c.KeepAlive(ctx, leaseId)

	for {
		select {
		case <-keepLiveCh:
			break
		case <-time.After(time.Duration(15) * time.Second):
			config.LOG.Error("A server lose heart", zap.Skip())
			return
		}
	}
}

// GetServerAddress get all addresses for this job
func GetServerAddress(job string) *Instance {
	client := newClient()
	resp, err := client.Get(context.Background(), fmt.Sprintf("%s/%s", consts.PrometheusJobKey, job), etcd.WithPrefix())
	if err != nil {
		config.LOG.Error("failed get server", zap.Skip())
		return nil
	}

	if resp.Count == 0 {
		return nil
	}
	addresses := make([]string, 0)
	for _, v := range resp.Kvs {
		addr := string(v.Value)
		if addr != "" {
			addresses = append(addresses, addr)
		}
	}
	conf := make([]*Conf, 1)
	conf[0] = &Conf{
		Targets: addresses,
		Labels: map[string]string{
			"job": job,
		}}
	return &Instance{
		Conf: conf,
	}
}

// GetAllServerAddress Get addresses for all the job
func GetAllServerAddress() []*Instance {
	service := config.CONFIG.Etcd.Services
	if len(service) == 0 {
		return nil
	}
	instances := make([]*Instance, len(service))
	for k, _ := range service {
		instances = append(instances, GetServerAddress(k))
	}
	return instances
}

// newClient return an etcd.Client
func newClient() *etcd.Client {
	client, err := etcd.New(etcd.Config{
		Endpoints:   []string{config.CONFIG.Etcd.Address},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		config.LOG.Error("", zap.Error(err))
	}
	return client
}
