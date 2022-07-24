package run

import (
	"context"
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

var CmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run app",
	Run:   runFunction,
}

var (
	host string
	port string
)

func init() {
	CmdRun.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "input hostname")
	CmdRun.Flags().StringVarP(&port, "port", "p", "8088", "input port")
}

func runFunction(cmd *cobra.Command, args []string) {

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

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
