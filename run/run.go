package run

import (
	"context"
	"fileserver/middleware/configx"
	"fileserver/middleware/logx"
	"fileserver/routers"
	"fileserver/utils"
	"fileserver/utils/iputil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func Run() {
	setting := configx.ServerSetting

	logx.InitLog(&setting.Log)

	utils.CheckFsHome()

	runtime.GOMAXPROCS(runtime.NumCPU())

	// 设置运行模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化路由
	router := routers.InitRouter()

	//addr := fmt.Sprintf("%s%s", ":", common.Port)
	addr := fmt.Sprintf("%s%s", ":", setting.System.Port)
	// 启动服务
	// 定义服务器
	httpServer := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    3600 * time.Second,
		WriteTimeout:   3600 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 利用 goroutine 启动监听
	go func() {
		// 服务连接
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.GetLogger().Sugar().Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("Local access address: http://127.0.0.1%s\n", addr)
	ips := iputil.GetIntranetIp()
	for _, ipStr := range ips {
		if ipStr != "" {
			fmt.Printf("Network access address: http://%s%s\n", ipStr, addr)
		}
	}
	// 优雅地重启或停止
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	// 反之，则执行后续代码
	<-quit
	logx.GetLogger().Sugar().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用 httpServer.Shutdown() 完成优雅停止
	// 调用时传递了一个上下文对象，对象中定义了超时时间
	if err := httpServer.Shutdown(ctx); err != nil {
		logx.GetLogger().Sugar().Fatal("Server Shutdown:", err)
	}
	logx.GetLogger().Sugar().Info("Server exiting")
}
