package routers

import (
	"fileserver/api"
	"fileserver/middleware/configx"
	"fileserver/middleware/cors"
	"fileserver/middleware/logx"
	"fileserver/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 全局中间件
	// 路由设置
	router := gin.New()

	router.Use(logx.ZapLogger(), logx.ZapRecovery(true))
	// 设置 Recovery 中间件，主要用于拦截 panic 错误，不至于导致进程崩掉
	router.Use(gin.Recovery())

	// 允许使用跨域请求  全局中间件
	router.Use(cors.Cors())
	router.GET("/ping", api.PingHandler)
	router.GET("/", api.IndexHandler)
	installPath, ok := utils.CheckFsHome()
	if ok {
		router.LoadHTMLGlob(fmt.Sprintf("%s%s", installPath, "/web/templates/*"))
		router.Static("/static", fmt.Sprintf("%s%s", installPath, "/web/static"))
	} else {
		logx.GetLogger().Sugar().Warn("Page not found!")
	}
	setting := configx.ServerSetting
	router.Static("/files", setting.System.StoragePath)
	router.GET("/list", api.List)
	router.POST("/upload", api.Upload)
	router.GET("/download/:fileName", api.Download)
	router.DELETE("/delete", api.Delete)
	router.NoRoute(api.NoRouteHandler)
	return router
}
