package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grm/global"
	"grm/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

type wsController struct {
}

var Wsc = wsController{}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func ReturnResp(data string, tag uint8, db uint8) []byte {
	return append([]byte{tag, db}, []byte(data)...)
}

func (con wsController) Ws(c *gin.Context) {

	type Cmd struct {
		Sk  string
		Db  int
		Cmd string
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()
	for {
	LOOP:
		mt, message, err := ws.ReadMessage()
		fmt.Println(message)
		if err != nil {
			err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, 0))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		var cmd Cmd
		json := json.NewDecoder(bytes.NewBuffer(message))
		err = json.Decode(&cmd)
		if err != nil {
			err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, uint8(cmd.Db)))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		var client *redis.Client

		client = global.GlobalClients[cmd.Sk]
		if client == nil {
			redisServer := global.GlobalConf.RedisServices[cmd.Sk]
			client, err = service.NewRedisClient(redisServer)
			if err != nil {
				err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, uint8(cmd.Db)))
				if err != nil {
					log.Println("write:", err)
					break
				}
				goto LOOP
			}
			global.GlobalClients[cmd.Sk] = client
			global.GlobalConf.RedisServices[cmd.Sk] = redisServer
		}
		err = client.Do(context.Background(), "select", cmd.Db).Err()
		if err != nil {
			err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, uint8(cmd.Db)))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		cmds := strings.Split(cmd.Cmd, " ")
		cmdL := make([]interface{}, len(cmds))
		for k, v := range cmds {

			if k == 0 && v == "select" {
				db, _ := strconv.Atoi(cmds[1])
				if cmd.Db != db {
					cmd.Db = db
				}
			}

			cmdL[k] = v
		}

		result, err := client.Do(context.Background(), cmdL...).Result()
		if err != nil {
			err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, uint8(cmd.Db)))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		var resultPut string
		switch val := result.(type) {
		case string:
			resultPut = val
		case int:
			resultPut = strconv.Itoa(val)
		case []interface{}:
			for _, v := range val {
				resultPut += fmt.Sprintf("%s \r\n", v.(string))
			}
		}

		err = ws.WriteMessage(mt, ReturnResp(resultPut, 1, uint8(cmd.Db)))
		if err != nil {
			log.Println("write:", err)
			break
		}
		goto LOOP
	}
}
