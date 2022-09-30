package run

import (
	"context"
	"fmt"
	"grm/common"
	"grm/global"
	"grm/glog"
	"grm/router"
	"grm/web"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"go.uber.org/zap"
)

type Services struct {
	Log service.Logger
	Srv *http.Server
	Cfg *service.Config
}

// 获取 service 对象
func GetSrv() service.Service {

	s := &Services{
		Cfg: &service.Config{
			Name:        "grm",
			DisplayName: "grm",
			Description: "Redis管理工具",
			Arguments:   []string{"run", "--daemon"},
		}}
	serv, err := service.New(s, s.Cfg)
	if err != nil {
		errMsg := fmt.Sprintf("%+v", errors.WithStack(err))
		glog.Logger.Error(errMsg, zap.Error(err))
		os.Exit(0)
	}
	s.Log, err = serv.SystemLogger(nil)
	if err != nil {
		errMsg := fmt.Sprintf("%+v", errors.WithStack(err))
		glog.Logger.Error(errMsg, zap.Error(err))
		os.Exit(0)
	}

	return serv
}

// 启动windows服务
func (srv *Services) Start(s service.Service) error {

	if srv.Log != nil {
		srv.Log.Info("Start run http server")
	}

	go srv.StarServer()

	return nil
}

// 停止windows服务
func (srv *Services) Stop(s service.Service) error {
	if srv.Log != nil {
		srv.Log.Info("Start stop http server")
	}
	return srv.Srv.Shutdown(context.Background())
}

// 运行web服务
func (srv *Services) StarServer() {

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	gin.DefaultWriter = io.Discard

	router := router.Init()
	router.StaticFS("/static", web.StaticsFs)

	srv.Srv = &http.Server{
		Addr:    net.JoinHostPort(global.GlobalConf.Host, global.GlobalConf.Port),
		Handler: router,
	}

	go func() {
		if err := srv.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errMsg := fmt.Sprintf("%+v", errors.WithStack(err))
			glog.Logger.Error(errMsg, zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	//输出LOGO
	common.ShowLogo(host, port)

	<-quit

	glog.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Srv.Shutdown(ctx); err != nil {
		errMsg := fmt.Sprintf("%+v", errors.WithStack(err))
		glog.Logger.Error(errMsg, zap.Error(err))
	}

	glog.Logger.Info("Server exiting")
}
