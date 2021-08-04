package client_conn

import (
	"context"
	"errors"
	"fmt"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/yongcycchen/mall-api/internal/config"
	"github.com/yongcycchen/mall-api/internal/service/slb"
	"github.com/yongcycchen/mall-api/internal/service/slb/etcdconfig"
	"github.com/yongcycchen/mall-api/pkg/util/grpc_interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	opts []grpc.DialOption
)

type Conn struct {
	ServerName string
	ServerPort string
}

func NewConn(serviceName string) (*Conn, error) {
	serviceNames := strings.Split(serviceName, "-")
	if len(serviceNames) < 1 {
		return nil, errors.New("NewConn.serviceNames is empty")
	}
	etcdServerUrls := config.GetEtcdV3ServerURLs()
	if len(etcdServerUrls) == 0 {
		return nil, fmt.Errorf("Can't not found env '%s'", config.ENV_ETCDV3_SERVER_URLS)
	}
	// load cache get client config
	currentConfig := loadClientConfig(serviceName)
	if currentConfig == nil {
		var err error
		serviceLB := slb.NewService(etcdServerUrls, serviceName)
		serviceConfig := etcdconfig.NewServiceConfig(serviceLB)
		currentConfig, err = serviceConfig.GetConfig()
		if err != nil {
			return nil, fmt.Errorf("serviceConfig.GetConfig err: %v", err)
		}
		storeClientConfig(serviceName, currentConfig)
	}

	return &Conn{
		ServerName: serviceName,
		ServerPort: currentConfig.ServicePort,
	}, nil
}

func (c *Conn) GetConn(ctx context.Context) (*grpc.ClientConn, error) {
	target := c.ServerName + ":" + c.ServerPort
	return grpc.DialContext(
		ctx,
		target,
		opts...,
	)
}

func init() {
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			grpc_interceptor.UnaryCtxHandleGRPC(),
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(2),
				grpc_retry.WithCodes(
					codes.Internal,
					codes.DeadlineExceeded,
				),
			),
		),
	))
	opts = append(opts, grpc.WithStreamInterceptor(
		grpc_middleware.ChainStreamClient(
			grpc_interceptor.StreamCtxHandleGRPC(),
		),
	))
}
