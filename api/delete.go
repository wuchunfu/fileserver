package api

import (
	"fileserver/middleware/configx"
	"fileserver/middleware/logx"
	"fileserver/utils/filex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type DeleteFiles struct {
	FileNameList []string `json:"fileNameList" binding:"required"`
}

// Delete 文件删除
func Delete(ctx *gin.Context) {
	setting := configx.ServerSetting
	storagePath := setting.System.StoragePath

	deleteFiles := &DeleteFiles{}
	ctx.Bind(deleteFiles)
	for _, fileName := range deleteFiles.FileNameList {
		storagePathAbs, _ := filepath.Abs(storagePath)
		isExistPath := filex.FilePathExists(storagePathAbs)
		if isExistPath {
			fullPath := fmt.Sprintf("%s/%s", storagePathAbs, fileName)
			// 删除文件
			err := os.Remove(fullPath)
			if err != nil {
				logx.GetLogger().Sugar().Errorf("Delete file failed!\n%s", err.Error())
				panic(err.Error())
			}
			logx.GetLogger().Sugar().Infof("File deleted successfully: %s", fileName)
		} else {
			logx.GetLogger().Sugar().Infof("File downloaded failed: %s", fileName)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "File deleted successfully!"})
}
