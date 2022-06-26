/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 10:08:32
 */
package global

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"grm/model"
	"path"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisServiceStorage map[string]RedisService
var HostName string
var Port string
var ConfigViper *viper.Viper
var Accounts map[string]string
var Limit int64

var UseClient GlobalClient

type GlobalClient struct {
	ConnectName string
	Db          int
	Client      *redis.Client
}

type RedisService struct {
	RedisService string
	Config       *redis.Options
	UseSsh       bool
	SSHConfig    model.SSHConfig
	Client       *redis.Client
}

func init() {
	RedisServiceStorage = make(map[string]RedisService)
	Accounts = make(map[string]string)

	//获取配置文件
	var configPath string
	flag.StringVar(&configPath, "c", "./config.yaml", "配置文件路径")
	flag.Parse()

	basePath := path.Base(configPath)
	fileInfo := strings.Split(basePath, ".")
	ConfigViper = viper.New()

	ConfigViper.AddConfigPath(path.Dir(configPath))
	ConfigViper.SetConfigName(fileInfo[0])
	ConfigViper.SetConfigType(fileInfo[1])
	err := ConfigViper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	HostName = ConfigViper.GetString("hostname")
	if HostName == "" {
		HostName = "127.0.0.1"
	}

	Port = ConfigViper.GetString("port")
	if Port == "" {
		Port = "8088"
	}

	Limit = ConfigViper.GetInt64("limit")
	if Limit == 0 {
		Limit = 100
	}

	connections := ConfigViper.Get("connections")

	if connections != nil {

		slice_conns := connections.([]interface{})

		for _, v := range slice_conns {

			vv := v.(map[string]interface{})

			optionConfig := &redis.Options{
				Addr:     vv["host"].(string) + ":" + vv["port"].(string),
				Password: vv["password"].(string),
				DB:       0,
			}

			// 生成md5 key
			m := md5.New()
			m.Write([]byte(vv["servicename"].(string)))
			key := hex.EncodeToString(m.Sum(nil))

			// if vv["usessh"].(bool) {
			// 	sshConfig := vv["sshconfig"].(map[string]interface{})
			// 	cli, err := common.GetSSHClient(sshConfig["sshusername"].(string), sshConfig["sshpassword"].(string), sshConfig["sshhost"].(string)+":"+sshConfig["sshport"].(string))
			// 	if nil != err {
			// 		panic(err)
			// 	}
			// 	optionConfig.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			// 		return cli.Dial(network, addr)
			// 	}
			// }

			// client := redis.NewClient(optionConfig)

			// client.AddHook(common.RedisLog{
			// 	Logger: common.NewLogger(vv["servicename"].(string)),
			// })

			// _, err := client.Ping(context.Background()).Result()
			// if err != nil {
			// 	panic(vv["servicename"].(string) + "连接失败:" + err.Error())
			// }

			RsSlice := RedisService{
				RedisService: vv["servicename"].(string),
				Config:       optionConfig,
			}

			RedisServiceStorage[key] = RsSlice

		}

	}

	accounts := ConfigViper.Get("accounts")
	if accounts != nil {
		slice_account := accounts.([]interface{})
		for _, account := range slice_account {
			accountMap := account.(map[string]interface{})
			Accounts[accountMap["account"].(string)] = accountMap["password"].(string)
		}
	}

}
