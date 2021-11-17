package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuchunfu/fileserver/middleware/logx"
	"github.com/wuchunfu/fileserver/utils/filex"
	"path/filepath"
)

// Download 文件下载
func Download(ctx *gin.Context) {
	dataMap := make(map[string]string)
	err := ctx.BindJSON(&dataMap)
	if err != nil {
		return
	}

	fileAbsPath := dataMap["filePath"]
	logx.GetLogger().Sugar().Infof("fileAbsPath: %s", fileAbsPath)

	fileName := filepath.Base(fileAbsPath)
	// 强制浏览器下载，设置浏览器是否为直接下载文件，且为浏览器指定下载文件的名字
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	// 浏览器下载或预览
	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", fileName))
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Type", "application/octet-stream; charset=utf-8;")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Expires", "0")
	// 如果缓存过期了，会再次和原来的服务器确定是否为最新数据，而不是和中间的proxy
	ctx.Header("Cache-Control", "must-revalidate")
	// 以流的形式下载必须设置这一项，否则前端下载下来的文件会出现格式不正确或已损坏的问题
	ctx.Header("Response-Type", "blob")
	ctx.Header("Pragma", "public")
	isExistPath := filex.FilePathExists(fileAbsPath)
	if isExistPath {
		ctx.File(fileAbsPath)
	} else {
		logx.GetLogger().Sugar().Infof("File downloaded failed: %s", fileName)
	}
	logx.GetLogger().Sugar().Infof("File downloaded successfully: %s", fileName)
}
