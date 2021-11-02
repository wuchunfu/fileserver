package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchunfu/fileserver/middleware/configx"
	"github.com/wuchunfu/fileserver/middleware/logx"
	"github.com/wuchunfu/fileserver/utils/bytex"
	"github.com/wuchunfu/fileserver/utils/filex"
	"net/http"
	"os"
	"path/filepath"
)

// Upload 文件上传
func Upload(ctx *gin.Context) {
	setting := configx.ServerSetting
	storagePath := setting.System.StoragePath

	storagePathAbs, _ := filepath.Abs(storagePath)
	form, _ := ctx.MultipartForm()
	if form != nil {
		isExistPath := filex.FilePathExists(storagePathAbs)
		if !isExistPath {
			filex.MkdirAll(storagePathAbs)
		}
		// 进入存储目录
		os.Chdir(storagePathAbs)
		// 获取所有上传文件信息
		files := form.File["formDataFile"]
		if len(files) <= 0 {
			ctx.JSON(http.StatusNoContent, gin.H{
				"code": http.StatusNoContent,
				"msg":  "The file is empty!",
				"data": nil,
			})
			return
		}
		// 循环对每个文件进行处理
		for _, file := range files {
			fileName := file.Filename
			fileSize := bytex.FormatFileSize(file.Size)
			err := ctx.SaveUploadedFile(file, fileName)
			if err != nil {
				logx.GetLogger().Sugar().Errorf("File upload failed!\n%s", err.Error())
				panic(err.Error())
			}
			logx.GetLogger().Sugar().Infof("File uploaded successfully: fileName：%s fileSize: %s\n", fileName, fileSize)
		}
		// 退出存储目录
		defer os.Chdir(storagePathAbs)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "File uploaded successfully!",
		"data": nil,
	})
}
