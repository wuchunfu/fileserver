package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchunfu/fileserver/middleware/configx"
	"github.com/wuchunfu/fileserver/utils/bytex"
	"github.com/wuchunfu/fileserver/utils/datetimex"
	"net/http"
	"os"
	"path/filepath"
)

type FileList struct {
	FileName string `json:"fileName"`
	FileSize string `json:"fileSize"`
	DateTime string `json:"dateTime"`
}

// List 文件列表
func List(ctx *gin.Context) {
	setting := configx.ServerSetting
	storagePath := setting.System.StoragePath

	fileList := make([]FileList, 0)
	storageAbs, _ := filepath.Abs(storagePath)
	// 遍历目录，读出文件名、大小
	filepath.Walk(storageAbs, func(filePath string, fileInfo os.FileInfo, err error) error {
		if nil == fileInfo {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}
		list := &FileList{
			FileName: fileInfo.Name(),
			FileSize: bytex.FormatFileSize(fileInfo.Size()),
			DateTime: datetimex.FormatDateTime(fileInfo.ModTime()),
		}
		fileList = append(fileList, *list)
		return nil
	})
	// 返回目录json数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Get data successfully!",
		"data": fileList,
	})
}
