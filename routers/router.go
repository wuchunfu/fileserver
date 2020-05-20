package routers

import (
	"fileserver/api"
	"fileserver/common"
	"fileserver/middleware/cors"
	"fileserver/middleware/logger"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 全局中间件
	//// 日志文件
	//fullPath := path.Join(common.LogPath, common.LogName)
	//// 记录到文件。
	//logFile, _ := os.Create(fullPath)
	//// 自定义日志格式
	//logConfig := gin.LoggerConfig{
	//	Formatter: func(params gin.LogFormatterParams) string {
	//		// 你的自定义格式
	//		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//			params.ClientIP,
	//			params.TimeStamp.Format(time.RFC1123),
	//			params.Method,
	//			params.Path,
	//			params.Request.Proto,
	//			params.StatusCode,
	//			params.Latency,
	//			params.Request.UserAgent(),
	//			params.ErrorMessage,
	//		)
	//	},
	//	Output: logFile,
	//}

	// 路由设置
	router := gin.New()

	router.Use(logger.WriteLogToFile())
	//router.Use(gin.LoggerWithConfig(logConfig))
	// 设置 Recovery 中间件，主要用于拦截 panic 错误，不至于导致进程崩掉
	router.Use(gin.Recovery())

	// 允许使用跨域请求  全局中间件
	router.Use(cors.Cors())
	router.GET("/ping", api.PingHandler)
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../web/templates/*")
	} else {
		router.LoadHTMLGlob("./web/templates/*")
	}
	router.GET("/", api.IndexHandler)
	router.Static("/static", "./web/static")
	router.Static("/files", common.StoragePath)
	router.GET("/list", api.List)
	router.POST("/upload", api.Upload)
	router.GET("/download/:fileName", api.Download)
	router.DELETE("/delete", api.Delete)
	router.NoRoute(api.NoRouteHandler)
	return router
}
