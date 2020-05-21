package api

import (
	"fileserver/common"
	"fileserver/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

// 文件上传
func Upload(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	if form != nil {
		utils.MkdirAll(common.StoragePath)
		currentDirectoryAbs, _ := filepath.Abs("./")
		storageAbs, _ := filepath.Abs(common.StoragePath)
		// 进入存储目录
		os.Chdir(storageAbs)
		// 退出存储目录
		defer os.Chdir(currentDirectoryAbs)
		// 获取所有上传文件信息
		files := form.File["formDataFile"]
		// 循环对每个文件进行处理
		for _, file := range files {
			fileName := file.Filename
			fileSize := utils.FormatFileSize(file.Size)
			err := ctx.SaveUploadedFile(file, fileName)
			if err != nil {
				logrus.Errorf("File upload failed!\n%s", err.Error())
				panic(err.Error())
			}
			logrus.Infof("File uploaded successfully: fileName：%s fileSize: %s\n", fileName, fileSize)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "File uploaded successfully!"})
}
