package api

import (
	"fileserver/middleware/configx"
	"fileserver/middleware/logx"
	"fileserver/utils"
	"fileserver/utils/filex"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// Upload 文件上传
func Upload(ctx *gin.Context) {
	setting := configx.ServerSetting
	storagePath := setting.System.StoragePath

	installPath, ok := utils.CheckFsHome()
	if ok {
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
			// 循环对每个文件进行处理
			for _, file := range files {
				fileName := file.Filename
				fileSize := utils.FormatFileSize(file.Size)
				err := ctx.SaveUploadedFile(file, fileName)
				if err != nil {
					logx.GetLogger().Sugar().Errorf("File upload failed!\n%s", err.Error())
					panic(err.Error())
				}
				logx.GetLogger().Sugar().Infof("File uploaded successfully: fileName：%s fileSize: %s\n", fileName, fileSize)
			}
			// 退出存储目录
			defer os.Chdir(installPath)
		}
		ctx.JSON(http.StatusOK, gin.H{"msg": "File uploaded successfully!"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"msg": "File uploaded failed!"})
	}
}
