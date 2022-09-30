package run

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grm/common"
	"grm/global"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var CmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run app",
	Run:   runFunction,
}

var (
	host      string
	port      string
	confirm   string
	stop      bool
	start     bool
	install   bool
	uninstall bool
	daemon    bool
)

func init() {
	confirm = "y"
	CmdRun.Flags().StringVarP(&host, "host", "H", global.GlobalConf.Host, "input hostname")
	CmdRun.Flags().StringVarP(&port, "port", "p", global.GlobalConf.Port, "input port")
	CmdRun.PersistentFlags().BoolVarP(&install, "install", "i", false, "install service")
	CmdRun.PersistentFlags().BoolVarP(&uninstall, "uninstall", "u", false, "uninstall service")
	CmdRun.PersistentFlags().BoolVarP(&start, "start", "", false, "start service")
	CmdRun.PersistentFlags().BoolVarP(&stop, "stop", "", false, "stop service")
	CmdRun.PersistentFlags().BoolVarP(&daemon, "daemon", "", false, "daemon service")
}

func runFunction(cmd *cobra.Command, args []string) {

	if !uninstall && !start && !stop && !daemon {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 33, "================", 0x1B)
		fmt.Printf("%c[%d;%d;%dm  Host:%s%c[0m \n", 0x1B, 0, 40, 33, host, 0x1B)
		fmt.Printf("%c[%d;%d;%dm  Port:%s%c[0m \n", 0x1B, 0, 40, 33, port, 0x1B)
		fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 33, "================", 0x1B)
		fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 33, "Do you want to run the app at this address? Y/N", 0x1B)
		fmt.Scan(&confirm)

		if strings.ToUpper(confirm) == "N" {
			os.Exit(0)
		}
	}

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

	s := GetSrv()

	if install || uninstall || start || stop {
		if install {
			if err := s.Install(); err != nil {
				fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 31, err.Error(), 0x1B)
				return
			}
			fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 32, "执行完成!", 0x1B)
			return
		}

		if uninstall {
			if err := s.Uninstall(); err != nil {
				fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 31, err.Error(), 0x1B)
				return
			}
			fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 32, "执行完成，如服务未卸载请先确保服务已停止运行!", 0x1B)
			return
		}

		if start {
			if err := s.Start(); err != nil {
				fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 31, err.Error(), 0x1B)
				return
			}
			fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 32, "执行完成，如服务不能正常访问请查看日志记录!", 0x1B)
			return
		}

		if stop {
			if err := s.Stop(); err != nil {
				fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 31, err.Error(), 0x1B)
				return
			}
			fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 32, "执行完成，如未停止服务请查看日志记录!", 0x1B)
			return
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m \n", 0x1B, 0, 40, 31, err.Error(), 0x1B)
	}

}
