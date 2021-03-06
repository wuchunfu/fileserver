package api

import (
	"fileserver/common"
	"fileserver/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type FileList struct {
	FileName string `json:"fileName"`
	FileSize string `json:"fileSize"`
	DateTime string `json:"dateTime"`
}

// 文件列表
func List(ctx *gin.Context) {
	fileList := make([]FileList, 0)
	storageAbs, _ := filepath.Abs(common.StoragePath)
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
			FileSize: utils.FormatFileSize(fileInfo.Size()),
			DateTime: utils.FormatDateTime(fileInfo.ModTime()),
		}
		fileList = append(fileList, *list)
		return nil
	})
	// 返回目录json数据
	ctx.JSON(http.StatusOK, gin.H{"data": fileList})
}
