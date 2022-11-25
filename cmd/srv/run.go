package srv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grm/common"
	"grm/global"
	"grm/router"
	"grm/web"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var cmdSrvRun = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Run:   srvRunFunc,
}

var (
	host string
	port string
)

func init() {
	cmdSrvRun.Flags().StringVarP(&host, "host", "H", global.GlobalConf.Host, "input hostname")
	cmdSrvRun.Flags().StringVarP(&port, "port", "p", global.GlobalConf.Port, "input port")
}

func srvRunFunc(cmd *cobra.Command, args []string) {

	storeHP()
	//开启服务
	start()
}

func storeHP() {
	//保存host和port信息
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	global.GlobalConf.Host = host
	global.GlobalConf.Port = port
	err := encoder.Encode(global.GlobalConf)

	if err != nil {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 31, err.Error(), 0x1B)
		os.Exit(0)
	}

	if err = common.WriteData(buffer.Bytes()); err != nil {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 31, err.Error(), 0x1B)
		os.Exit(0)
	}

}

func start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	router := router.Init()

	router.StaticFS("/static", web.StaticsFs)

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
