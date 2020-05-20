package api

import (
	"fileserver/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

type DeleteFiles struct {
	FileNameList []string `json:"fileNameList" binding:"required"`
}

// 文件删除
func Delete(ctx *gin.Context) {
	deleteFiles := &DeleteFiles{}
	ctx.Bind(deleteFiles)
	for _, file := range deleteFiles.FileNameList {
		storageAbs, _ := filepath.Abs(common.StoragePath)
		fullPath := storageAbs + "/" + file
		// 删除文件
		err := os.Remove(fullPath)
		if err != nil {
			logrus.Errorf("Delete file failed!\n%s", err.Error())
			panic(err.Error())
		}
		logrus.Infof("File deleted successfully: %s", file)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "File deleted successfully!"})
}
