package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "文件共享系统",
	})
}

func NoRouteHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{
		"status": "404",
		"msg":    "Sorry, the page you visited does not exist.",
	})
}
