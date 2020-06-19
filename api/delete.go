package api

import (
	"fileserver/common"
	"fileserver/utils"
	"fmt"
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
	for _, fileName := range deleteFiles.FileNameList {
		storagePathAbs, _ := filepath.Abs(common.StoragePath)
		isExistPath := utils.IsExistPath(storagePathAbs)
		if isExistPath {
			fullPath := fmt.Sprintf("%s/%s", storagePathAbs, fileName)
			// 删除文件
			err := os.Remove(fullPath)
			if err != nil {
				logrus.Errorf("Delete file failed!\n%s", err.Error())
				panic(err.Error())
			}
			logrus.Infof("File deleted successfully: %s", fileName)
		} else {
			logrus.Infof("File downloaded failed: %s", fileName)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "File deleted successfully!"})
}
