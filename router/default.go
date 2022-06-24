/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:17:35
 */
package router

import (
	"grm/controllers"
	"grm/global"
	"grm/middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("goredismanagerphper"))
	router.Use(middleware.StaticCache(), gzip.Gzip(gzip.DefaultCompression), sessions.Sessions("goredismanager", store))

	app := router.Group("/")

	if len(global.Accounts) > 0 {
		//用户验证
		app.Use(middleware.UserAuth())
	}

	connRouter := app.Group("/conn")
	{
		connRouter.GET("/list", controllers.Cc.List)
	}

	return router
}
