package run

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
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run app",
	Run:   runFunction,
}

var (
	host    string
	port    string
	confirm string
)

func init() {
	confirm = "n"
	CmdRun.Flags().StringVarP(&host, "host", "H", global.GlobalConf.Host, "input hostname")
	CmdRun.Flags().StringVarP(&port, "port", "p", global.GlobalConf.Port, "input port")
}

func runFunction(cmd *cobra.Command, args []string) {

	fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 33, "================", 0x1B)
	fmt.Printf("%c[%d;%d;%dm  Host:%s%c[0m \n", 0x1B, 0, 40, 33, host, 0x1B)
	fmt.Printf("%c[%d;%d;%dm  Port:%s%c[0m \n", 0x1B, 0, 40, 33, port, 0x1B)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 33, "================", 0x1B)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 33, "Do you want to run the app at this address? Y/N", 0x1B)
	fmt.Scan(&confirm)

	if strings.ToUpper(confirm) == "N" {
		os.Exit(0)
	}

	//保存host和port信息
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	global.GlobalConf.Host = host
	global.GlobalConf.Port = port
	err := encoder.Encode(global.GlobalConf)
	if err != nil {
		fmt.Println(err)
	}
	if err = common.WriteData(buffer.Bytes()); err != nil {
		fmt.Println(err)
	}

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
