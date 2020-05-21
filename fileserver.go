package main

import (
	"context"
	"fileserver/common"
	"fileserver/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 设置运行模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化路由
	router := routers.InitRouter()

	// 启动服务
	// 定义服务器
	httpServer := &http.Server{
		Addr:           common.Addr,
		Handler:        router,
		ReadTimeout:    3600 * time.Second,
		WriteTimeout:   3600 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 利用 goroutine 启动监听
	go func() {
		// 服务连接
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Errorf("Failed to get hostname!\n%s", err.Error())
	}
	fmt.Printf("Service listening address: http://%s%s\n", hostname, common.Addr)
	// 优雅地重启或停止
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	// 反之，则执行后续代码
	<-quit
	logrus.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用 httpServer.Shutdown() 完成优雅停止
	// 调用时传递了一个上下文对象，对象中定义了超时时间
	if err := httpServer.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Info("Server exiting")
}
