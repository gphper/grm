/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 10:08:32
 */
package global

import (
	"bytes"
	"encoding/json"
	"fmt"
	"grm/common"
	"grm/model"
	"io"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

var globalClients sync.Map

func SetClient(key string, client *redis.Client) {
	globalClients.Store(key, client)
}

func GetClient(key string) *redis.Client {
	client, ok := globalClients.Load(key)
	if !ok {
		return nil
	}
	return client.(*redis.Client)
}

type GlobalConfig struct {
	Accounts        map[string]string
	RedisServicesCp map[string]RedisService
	RedisServices   sync.Map
	Separator       string
	Host            string
	Port            string
	Tree            bool
}

type RedisService struct {
	RedisService string
	Config       model.RedisConfig
	UseSsh       bool
	SSHConfig    model.SSHConfig
}

var GlobalConf = GlobalConfig{
	Accounts:      make(map[string]string),
	RedisServices: sync.Map{},
	Separator:     ":",
	Host:          "0.0.0.0",
	Port:          "8088",
	Tree:          true,
}

func init() {

	data, err := common.ReadData()
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(0)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(data))
	err = decoder.Decode(&GlobalConf)

	for k, v := range GlobalConf.RedisServicesCp {
		GlobalConf.RedisServices.Store(k, v)
	}

	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(0)
	}
}
