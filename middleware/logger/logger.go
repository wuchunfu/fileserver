package logger

import (
	"fileserver/common"
	"fileserver/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

// 日志记录到文件
func WriteLogToFile() gin.HandlerFunc {
	// 如果文件夹不存在就创建
	utils.MkdirAll(common.LogPath)
	// 日志文件
	fullPath := path.Join(common.LogPath, common.LogName)
	// 写入文件
	filePath, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logrus.Errorf("fail to write to file\n%s", err.Error())
	}
	// 实例化
	loggers := logrus.New()
	// 设置输出
	loggers.SetOutput(filePath)
	//设置日志格式
	//loggers.SetFormatter(&logrus.TextFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	loggers.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	level := common.LogLevel
	switch strings.ToLower(level) {
	case "debug":
		loggers.SetLevel(logrus.DebugLevel)
	case "info":
		loggers.SetLevel(logrus.InfoLevel)
	case "warn":
		loggers.SetLevel(logrus.WarnLevel)
	case "error":
		loggers.SetLevel(logrus.ErrorLevel)
	default:
		loggers.SetLevel(logrus.InfoLevel)
	}
	return func(ctx *gin.Context) {
		// 请求客户端 ip
		clientIP := ctx.ClientIP()
		// 开始时间
		startTime := time.Now()
		// 处理请求
		ctx.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求时间
		requestTime := startTime.Format("2006-01-02 15:04:05")
		// 请求方式
		requestMethod := ctx.Request.Method
		// 请求协议
		requestProto := ctx.Request.Proto
		// 请求路由
		requestUri := ctx.Request.RequestURI
		// 状态码
		statusCode := ctx.Writer.Status()
		// 用户代理
		userAgent := ctx.Request.UserAgent()
		// 日志格式
		loggers.WithFields(logrus.Fields{
			"clientIP":      clientIP,
			"requestTime":   requestTime,
			"requestMethod": requestMethod,
			"requestUri":    requestUri,
			"requestProto":  requestProto,
			"statusCode":    statusCode,
			"latencyTime":   latencyTime,
			"userAgent":     userAgent,
		}).Log(loggers.GetLevel())
		//loggers.Logf(loggers.GetLevel(), fmt.Sprintf("%s - [%s] %s %s %s %d %s [%s]",
		//	clientIP,
		//	requestTime,
		//	requestMethod,
		//	requestUri,
		//	requestProto,
		//	statusCode,
		//	latencyTime,
		//	userAgent,
		//))
	}
}
