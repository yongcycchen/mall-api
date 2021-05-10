package vars

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Version    = "1.0.0"
	AppTypeWeb = 1
)

type Application struct {
	Name       string
	Type       int32
	LoadConfig func() error
	SetupVars  func() error
	StopFunc   func() error
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
