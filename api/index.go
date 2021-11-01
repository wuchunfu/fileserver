package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "pong",
		"data": nil,
	})
}

func RedirectIndex(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/ui")
}

func NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code": http.StatusNotFound,
		"msg":  "Sorry, the page you visited does not exist.",
		"data": nil,
	})
}
