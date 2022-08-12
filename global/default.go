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

	"github.com/go-redis/redis"
)

var GlobalClients map[string]*redis.Client

type GlobalConfig struct {
	Accounts      map[string]string
	RedisServices map[string]RedisService
	Separator     string
	Tree          bool
}

type RedisService struct {
	RedisService string
	Config       model.RedisConfig
	UseSsh       bool
	SSHConfig    model.SSHConfig
}

var GlobalConf = GlobalConfig{
	Accounts:      make(map[string]string),
	RedisServices: make(map[string]RedisService),
	Separator:     ":",
	Tree:          true,
}

func init() {

	GlobalClients = make(map[string]*redis.Client)

	data, err := common.ReadData()
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(0)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(data))
	err = decoder.Decode(&GlobalConf)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(0)
	}
}
