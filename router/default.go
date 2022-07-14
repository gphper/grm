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

	app := router.Group("/api")

	if len(global.Accounts) > 0 {
		//用户验证
		app.Use(middleware.UserAuth())
	}

	connRouter := app.Group("/conn")
	{
		connRouter.GET("/list", controllers.Cc.List)
		connRouter.POST("/add", controllers.Cc.Add)
		connRouter.POST("/test", controllers.Cc.TestConn)
	}

	indexRouter := app.Group("/index")
	{
		indexRouter.POST("/open", controllers.Ic.Open)
		indexRouter.POST("/getkeys", controllers.Ic.GetKeys)
		indexRouter.POST("/getkeytype", controllers.Ic.GetKeyType)
		indexRouter.POST("/delkey", controllers.Ic.DelKey)
		indexRouter.POST("/ttlkey", controllers.Ic.TtlKey)
		indexRouter.POST("/serinfo", controllers.Ic.SerInfo)
	}

	stringRouter := app.Group("/string")
	{
		stringRouter.POST("/show", controllers.Sc.Show)
	}

	listRouter := app.Group("/list")
	{
		listRouter.POST("/show", controllers.Lc.Show)
		listRouter.POST("/del", controllers.Lc.Del)
		listRouter.POST("/add", controllers.Lc.AddItem)
	}

	hashRouter := app.Group("/hash")
	{
		hashRouter.POST("/show", controllers.Hc.Show)
		hashRouter.POST("/del", controllers.Hc.Del)
		hashRouter.POST("/add", controllers.Hc.AddItem)
	}

	setRouter := app.Group("/set")
	{
		setRouter.POST("/show", controllers.Setc.Show)
		setRouter.POST("/del", controllers.Setc.Del)
		setRouter.POST("/add", controllers.Setc.AddItem)
	}

	zsetRouter := app.Group("/zset")
	{
		zsetRouter.POST("/show", controllers.Zc.Show)
		zsetRouter.POST("/del", controllers.Zc.Del)
		zsetRouter.POST("/add", controllers.Zc.AddItem)
	}

	streamRouter := app.Group("/stream")
	{
		streamRouter.POST("/show", controllers.Stc.Show)
		streamRouter.POST("/del", controllers.Stc.Del)
		streamRouter.POST("/add", controllers.Stc.AddItem)
	}

	return router
}
