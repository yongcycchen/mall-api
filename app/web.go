package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/cloudflare/tableflip"
	"github.com/robfig/cron/v3"
	"github.com/yongcycchen/mall-api/internal/config"
	"github.com/yongcycchen/mall-api/internal/logging"
	"github.com/yongcycchen/mall-api/vars"
)

const localAddr = "0.0.0.0:"

func RunApplication(application *vars.WEBApplication) {
	if application.Name == "" {
		// logging.FATAL("Application name can't be empty")
		logging.Fatal("Application name can't not be empty")
	}
	application.Type = vars.AppTypeWeb
	vars.App = application
	err := runApp(application)
	if err != nil {
		logging.Fatalf("App.runApp err: %v", err)
	}
}

// setupGRPCVars ...
func setupWEBVars(webApp *vars.WEBApplication) error {
	err := setupCommonVars(webApp)
	if err != nil {
		return err
	}

	return nil
}

func runApp(webApp *vars.WEBApplication) error {
	// load config
	err := config.LoadDefaultConfig(webApp.Application)
	if err != nil {
		return err
	}
	if webApp.LoadConfig != nil {
		err = webApp.LoadConfig()
		if err != nil {
			return err
		}
	}

	// 2. init application
	err = initApplication(webApp.Application)
	if err != nil {
		return err
	}

	// 3. setup vars
	err = setupWEBVars(webApp)
	if err != nil {
		return err
	}
	if webApp.SetupVars != nil {
		err = webApp.SetupVars()
		if err != nil {
			return fmt.Errorf("App.SetupVars err: %v", err)
		}
	}
	// 4. setup server monitor

	// 5 run task
	if webApp.RegisterTasks != nil {
		cronTasks := webApp.RegisterTasks()
		if len(cronTasks) != 0 {
			cn := cron.New(cron.WithSeconds())
			for i := 0; i < len(cronTasks); i++ {
				if cronTasks[i].TaskFunc != nil {
					_, err = cn.AddFunc(cronTasks[i].Cron, cronTasks[i].TaskFunc)
					if err != nil {
						logging.Fatalf("App run cron task err: %v", err)
					}
				}
			}
			cn.Start()
			logging.Info("App run cron task")
		}
	}

	// 6. set init service port
	var addr string
	if webApp.EndPort != 0 {
		addr = localAddr + strconv.Itoa(webApp.EndPort)
	} else if vars.ServerSetting.EndPort != 0 {
		addr = localAddr + strconv.Itoa(vars.ServerSetting.EndPort)
	}

	// 7. run http server

	if webApp.RegisterHttpRoute == nil {
		logging.Fatalf("App RegisterHttpRoute nil ??")
	}
	wd, _ := os.Getwd()
	pidFile := fmt.Sprintf("%s/%s.pid", wd, webApp.Name)
	if vars.ServerSetting.PIDFile != "" {
		pidFile = vars.ServerSetting.PIDFile
	}
	kp := new(KProcess)
	network := "tcp"
	if vars.ServerSetting != nil && vars.ServerSetting.Network != "" {
		network = vars.ServerSetting.Network
	}
	ln, err := kp.Listen(network, addr, pidFile)
	if err != nil {
		logging.Fatalf("App kprocess listen err: %v", err)
	}
	ginEngine := webApp.RegisterHttpRoute()
	serve := http.Server{
		Handler:      ginEngine,
		ReadTimeout:  time.Duration(vars.ServerSetting.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(vars.ServerSetting.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(vars.ServerSetting.IdleTimeout) * time.Second,
	}
	go func() {
		err = serve.Serve(ln)
		if err != nil {
			logging.Fatalf("App run Serve err: %v", err)

		}
	}()
	<-kp.Exit()

	appPrepareForceExit()
	err = serve.Shutdown(context.Background())
	if err != nil {
		logging.Fatalf("App server Shutdown err: %v", err)
	}
	err = appShutdown(webApp.Application)

	return err
}

type KProcess struct {
	pidFile   string
	pid       int
	processUp *tableflip.Upgrader
}

// This shows how to use the upgrader
// with the graceful shutdown facilities of net/http.
func (k *KProcess) Listen(network, addr, pidFile string) (ln net.Listener, err error) {
	k.pid = os.Getpid()
	logging.Infof(fmt.Sprintf("exec process pid %d \n", k.pid))

	k.processUp, err = tableflip.New(tableflip.Options{
		UpgradeTimeout: 500 * time.Millisecond,
		PIDFile:        pidFile,
	})
	if err != nil {
		return nil, err
	}
	k.pidFile = pidFile

	go k.signal(k.upgrade, k.stop)

	// Listen must be called before Ready
	if network != "" && addr != "" {
		ln, err = k.processUp.Listen(network, addr)
		if err != nil {
			return nil, err
		}
	}
	if err := k.processUp.Ready(); err != nil {
		return nil, err
	}

	return ln, nil
}

func (k *KProcess) stop() error {
	if k.processUp != nil {
		k.processUp.Stop()
		return os.Remove(k.pidFile)
	}
	return nil
}

func (k *KProcess) upgrade() error {
	if k.processUp != nil {
		return k.processUp.Upgrade()
	}
	return nil
}

func (k *KProcess) Exit() <-chan struct{} {
	if k.processUp != nil {
		return k.processUp.Exit()
	}
	ch := make(chan struct{})
	close(ch)
	return ch
}

func (k *KProcess) signal(upgradeFunc, stopFunc func() error) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)
	for s := range sig {
		switch s {
		case syscall.SIGTERM:
			if stopFunc != nil {
				err := stopFunc()
				if err != nil {
					logging.Infof("KProcess exec stopFunc failed:%v\n", err)
				}
				logging.Infof("process [%d] stop...\n", k.pid)
			}
			return
		case syscall.SIGUSR1, syscall.SIGUSR2:
			if upgradeFunc != nil {
				err := upgradeFunc()
				if err != nil {
					logging.Infof("KProcess exec Upgrade failed:%v\n", err)
				}
				logging.Infof("process [%d] restart...\n", k.pid)
			}
		}
	}
}
