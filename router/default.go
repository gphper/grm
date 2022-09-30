/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-19 20:17:35
 */
package router

import (
	"grm/controllers"
	"grm/glog"
	"grm/middleware"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()

	store := cookie.NewStore([]byte("goredismanagerphper"))
	router.Use(middleware.StaticCache(), gzip.Gzip(gzip.DefaultCompression), sessions.Sessions("goredismanager", store))
	router.Use(gin.Logger(), middleware.GinRecovery(glog.NewLogger("gin_error.log"), true))
	router.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/static/#")
	})

	app := router.Group("/grmapix")

	userRouter := app.Group("/user")
	{
		userRouter.POST("/login", controllers.Loginc.Login)
	}

	connRouter := app.Group("/conn")
	connRouter.Use(middleware.UserAuth())
	{
		connRouter.GET("/list", controllers.Cc.List)
		connRouter.POST("/add", controllers.Cc.Add)
		connRouter.POST("/del", controllers.Cc.Del)
		connRouter.POST("/test", controllers.Cc.TestConn)
	}

	indexRouter := app.Group("/index")
	indexRouter.Use(middleware.UserAuth())
	{
		indexRouter.POST("/open", controllers.Ic.Open)
		indexRouter.POST("/getkeys", controllers.Ic.GetKeys)
		indexRouter.POST("/getkeytype", controllers.Ic.GetKeyType)
		indexRouter.POST("/delkey", controllers.Ic.DelKey)
		indexRouter.POST("/ttlkey", controllers.Ic.TtlKey)
		indexRouter.POST("/serinfo", controllers.Ic.SerInfo)
		indexRouter.POST("/luarun", controllers.Ic.LuaRun)
		indexRouter.Any("/setting", controllers.Ic.Setting)
	}

	stringRouter := app.Group("/string")
	stringRouter.Use(middleware.UserAuth())
	{
		stringRouter.POST("/show", controllers.Sc.Show)
		stringRouter.POST("/add", controllers.Sc.Add)
	}

	listRouter := app.Group("/list")
	listRouter.Use(middleware.UserAuth())
	{
		listRouter.POST("/show", controllers.Lc.Show)
		listRouter.POST("/del", controllers.Lc.Del)
		listRouter.POST("/add", controllers.Lc.AddItem)
	}

	hashRouter := app.Group("/hash")
	hashRouter.Use(middleware.UserAuth())
	{
		hashRouter.POST("/show", controllers.Hc.Show)
		hashRouter.POST("/del", controllers.Hc.Del)
		hashRouter.POST("/add", controllers.Hc.AddItem)
	}

	setRouter := app.Group("/set")
	setRouter.Use(middleware.UserAuth())
	{
		setRouter.POST("/show", controllers.Setc.Show)
		setRouter.POST("/del", controllers.Setc.Del)
		setRouter.POST("/add", controllers.Setc.AddItem)
	}

	zsetRouter := app.Group("/zset")
	zsetRouter.Use(middleware.UserAuth())
	{
		zsetRouter.POST("/show", controllers.Zc.Show)
		zsetRouter.POST("/del", controllers.Zc.Del)
		zsetRouter.POST("/add", controllers.Zc.AddItem)
	}

	streamRouter := app.Group("/stream")
	streamRouter.Use(middleware.UserAuth())
	{
		streamRouter.POST("/show", controllers.Stc.Show)
		streamRouter.POST("/del", controllers.Stc.Del)
		streamRouter.POST("/add", controllers.Stc.AddItem)
	}

	wsapp := app.Group("/ws")
	{
		wsapp.GET("/cmd", controllers.Wsc.Ws)
	}

	return router
}
