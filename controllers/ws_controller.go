package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grm/common"
	"grm/global"
	"grm/service"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
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
		Jwt string
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
		if err != nil {
			break
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

		//验证用户
		info, err := jwt.ParseWithClaims(cmd.Jwt, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("TyPbWNRjho"), nil
		})
		if err != nil {
			err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, uint8(cmd.Db)))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		uInfo := info.Claims.(jwt.MapClaims)
		expire, err := time.ParseInLocation("2006-01-02 15:04:05", uInfo["expire"].(string), time.Local)
		if err != nil {
			err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, uint8(cmd.Db)))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		if time.Until(expire).Seconds() < 0 {
			err = ws.WriteMessage(mt, ReturnResp("Jwt Token Expired", 0, uint8(cmd.Db)))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		var username string = uInfo["user"].(string)
		//判断是不是心跳信息
		if cmd.Sk == "ping" {
			err = ws.WriteMessage(mt, ReturnResp("ping", 1, uint8(cmd.Db)))
			if err != nil {
				log.Println("write:", err)
				break
			}
			goto LOOP
		}

		ctx := context.WithValue(context.Background(), "username", username)
		var client *redis.Client

		client = global.GetClient(cmd.Sk)
		if client == nil {
			redisServer, _ := global.GlobalConf.RedisServices.Load(cmd.Sk)
			client, err = service.NewRedisClient(redisServer.(global.RedisService))
			if err != nil {
				err = ws.WriteMessage(mt, ReturnResp(err.Error(), 0, uint8(cmd.Db)))
				if err != nil {
					log.Println("write:", err)
					break
				}
				goto LOOP
			}

			global.SetClient(cmd.Sk, client)
			global.GlobalConf.RedisServices.Store(cmd.Sk, redisServer)
		}
		err = client.Do(ctx, "select", cmd.Db).Err()
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

		result, err := client.Do(ctx, cmdL...).Result()
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
		case []interface{}:
			for _, v := range val {
				resultPut += fmt.Sprintf("%s \r\n", common.InterfaceToString(v))
			}
		default:
			resultPut = common.InterfaceToString(val)
		}

		err = ws.WriteMessage(mt, ReturnResp(resultPut, 1, uint8(cmd.Db)))
		if err != nil {
			log.Println("write:", err)
			break
		}
		goto LOOP
	}
}
