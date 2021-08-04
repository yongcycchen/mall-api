package client_conn

import (
	"context"
	"math"
	"time"

	"github.com/bluele/gcache"
	"github.com/yongcycchen/mall-api/internal/service/slb/etcdconfig"
	"github.com/yongcycchen/mall-api/vars"
)

var (
	clientConfigCache gcache.Cache // 客户端缓存
	ctx               = context.Background()
)

func init() {
	clientConfigCache = gcache.New(math.MaxInt8).LRU().Build()
}

func storeClientConfig(serviceName string, config *etcdconfig.Config) {
	err := clientConfigCache.SetWithExpire(serviceName, *config, 5*time.Minute)
	if err != nil {
		vars.FrameworkLogger.Errorf(ctx, "[kelvins] storeClientConfig err: %v, serviceName: %v,config: %+v", err, serviceName, config)
		return
	}
}

func loadClientConfig(serviceName string) *etcdconfig.Config {
	exist := clientConfigCache.Has(serviceName)
	if exist {
		obj, err := clientConfigCache.Get(serviceName)
		if err == nil && obj != nil {
			c, ok := obj.(etcdconfig.Config)
			if ok {
				return &c
			}
		}
		return nil
	}
	return nil
}
