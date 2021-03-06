package api

import (
	"fileserver/common"
	"fileserver/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

// 文件下载
func Download(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	// 设置浏览器是否为直接下载文件，且为浏览器指定下载文件的名字
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Expires", "0")
	// 如果缓存过期了，会再次和原来的服务器确定是否为最新数据，而不是和中间的proxy
	ctx.Header("Cache-Control", "must-revalidate")
	ctx.Header("Pragma", "public")
	storagePathAbs, _ := filepath.Abs(common.StoragePath)
	isExistPath := utils.IsExistPath(storagePathAbs)
	if isExistPath {
		fullPath := fmt.Sprintf("%s/%s", storagePathAbs, fileName)
		ctx.File(fullPath)
	} else {
		logrus.Infof("File downloaded failed: %s", fileName)
	}
	logrus.Infof("File downloaded successfully: %s", fileName)
}
