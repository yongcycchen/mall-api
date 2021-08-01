package vars

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/yongcycchen/mall-api/common/event"
	"github.com/yongcycchen/mall-api/common/log"
	"google.golang.org/grpc"
)

const (
	Version      = "1.0.0"
	AppTypeWeb   = 0
	AppTypeGrpc  = 1
	AppTypeCron  = 2
	AppTypeQueue = 3
	AppTypeHttp  = 4
)

type Application struct {
	Name           string
	Type           int32
	LoggerRootPath string
	LoggerLevel    string
	LoadConfig     func() error
	SetupVars      func() error
	StopFunc       func() error
}

type WEBApplication struct {
	*Application
	EndPort        int
	MonitorEndPort int
	// Monitor
	Mux *http.ServeMux
	// RegisterHttpRoute
	RegisterHttpRoute func() *gin.Engine
	// Tasks
	RegisterTasks func() []CronTask
}

// GRPCApplication ...
type GRPCApplication struct {
	*Application
	Port                    int64
	GRPCServer              *grpc.Server
	GatewayServeMux         *runtime.ServeMux
	Mux                     *http.ServeMux
	HttpServer              *http.Server
	TlsConfig               *tls.Config
	GKelvinsLogger          *log.LoggerContext
	GSysErrLogger           *log.LoggerContext
	UnaryServerInterceptors []grpc.UnaryServerInterceptor
	ServerOptions           []grpc.ServerOption
	RegisterGRPCServer      func(*grpc.Server) error
	RegisterGateway         func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
	RegisterHttpRoute       func(*http.ServeMux) error
	EventServer             *event.EventServer
	RegisterEventProducer   func(event.ProducerIface) error
}
type ListenerApplication struct {
	*Application
	EndPort        int
	MonitorEndPort int
	Network        string
	ReadTimeout    int
	WriteTimeout   int
	// Monitor
	Mux *http.ServeMux
	// EventHandler After Listener Server Accept
	EventHandler func(context.Context, []byte) ([]byte, error)
}
