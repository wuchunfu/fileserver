package cors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 请求头部
		origin := ctx.Request.Header.Get("Origin")
		// 声明请求头keys
		var headerKeys []string
		for k, _ := range ctx.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("Access-Control-Allow-Origin, Access-Control-Allow-Headers, %s", headerStr)
		} else {
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Headers"
		}
		if origin != "" {
			// 这是允许访问所有域
			ctx.Header("Access-Control-Allow-Origin", "*")
			// 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH, CONNECT, HEAD, OPTIONS, TRACE")
			// header的类型
			//ctx.Header("Access-Control-Allow-Headers", headerStr)
			ctx.Header("Access-Control-Allow-Headers", "*")
			// 允许跨域设置, 可以返回其他子段
			// 跨域关键设置 让浏览器可以解析
			//ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma, FooBar")
			ctx.Header("Access-Control-Expose-Headers", "*")
			// 缓存请求信息 单位为秒 (一天)
			ctx.Header("Access-Control-Max-Age", "86400")
			// 跨域请求是否需要带cookie信息 默认设置为true
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		// 请求方法
		method := ctx.Request.Method
		// 放行所有OPTIONS方法
		if method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			// 处理请求
			ctx.Next()
		}
	}
}
