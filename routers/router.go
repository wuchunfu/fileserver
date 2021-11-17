package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchunfu/fileserver/api"
	"github.com/wuchunfu/fileserver/middleware/configx"
	"github.com/wuchunfu/fileserver/middleware/cors"
	"github.com/wuchunfu/fileserver/middleware/logx"
	"github.com/wuchunfu/fileserver/web"
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
	router.GET("/", api.RedirectIndex)
	router.StaticFS("/ui", web.GetFS())
	//installPath, ok := utils.CheckFsHome()
	//if ok {
	//	router.LoadHTMLGlob(fmt.Sprintf("%s%s", installPath, "/web/templates/*"))
	//	router.Static("/static", fmt.Sprintf("%s%s", installPath, "/web/static"))
	//} else {
	//	logx.GetLogger().Sugar().Warn("Page not found!")
	//}
	setting := configx.ServerSetting
	router.Static("/files", setting.System.StoragePath)
	router.GET("/list", api.List)
	router.POST("/upload", api.Upload)
	router.POST("/download", api.Download)
	router.DELETE("/delete", api.Delete)
	router.GET("/changeFolder", api.ChangeFolder)
	router.POST("/createFolder", api.CreateFolder)
	router.NoRoute(api.NoRouteHandler)
	return router
}
