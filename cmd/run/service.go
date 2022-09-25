package run

import (
	"context"
	"fmt"
	"grm/common"
	"grm/global"
	"grm/router"
	"grm/web"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
)

type Services struct {
	Log service.Logger
	Srv *http.Server
	Cfg *service.Config
}

// 获取 service 对象
func GetSrv() service.Service {

	path, err := common.RootPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	File, err := common.OpenFile(path + "/log/service.log")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer File.Close()

	log.SetOutput(File)

	s := &Services{
		Cfg: &service.Config{
			Name:        "grm",
			DisplayName: "grm",
			Description: "Redis管理工具",
			Arguments:   []string{"run", "--daemon"},
		}}
	serv, err := service.New(s, s.Cfg)
	if err != nil {
		log.Printf("Set logger error:%s\n", err.Error())
	}
	s.Log, err = serv.SystemLogger(nil)
	if err != nil {
		log.Printf("Set logger error:%s\n", err.Error())
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
	log.Println("Server exiting")
	return srv.Srv.Shutdown(context.Background())
}

// 运行web服务
func (srv *Services) StarServer() {

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	path, err := common.RootPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 创建记录日志的文件
	f, err := common.OpenFile(path + "/log/grm_error.log")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	gin.DefaultErrorWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.Discard

	router := router.Init()

	router.StaticFS("/static", web.StaticsFs)

	srv.Srv = &http.Server{
		Addr:    net.JoinHostPort(global.GlobalConf.Host, global.GlobalConf.Port),
		Handler: router,
	}

	go func() {
		if err := srv.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	//输出LOGO
	common.ShowLogo(host, port)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
