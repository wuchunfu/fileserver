package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchunfu/fileserver/middleware/configx"
	"github.com/wuchunfu/fileserver/middleware/logx"
	"github.com/wuchunfu/fileserver/utils/bytex"
	"github.com/wuchunfu/fileserver/utils/datetimex"
	"github.com/wuchunfu/fileserver/utils/filex"
	"io/fs"
	"net/http"
	"path/filepath"
)

type FileList struct {
	BasePath string `json:"basePath"`
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FileSize string `json:"fileSize"`
	DateTime string `json:"dateTime"`
}

// List 文件列表
func List(ctx *gin.Context) {
	setting := configx.ServerSetting
	storagePath := setting.System.StoragePath

	fileList := make([]FileList, 0)
	storageAbsPath, _ := filepath.Abs(storagePath)
	isExistPath := filex.FilePathExists(storageAbsPath)
	if !isExistPath {
		filex.MkdirAll(storageAbsPath)
	}
	// 遍历目录，读出文件名、大小
	err := filepath.WalkDir(storageAbsPath, func(filePath string, info fs.DirEntry, err error) error {
		fileInfo, err := info.Info()
		if nil == fileInfo {
			return err
		}
		if filePath == storageAbsPath {
			return nil
		}
		list := &FileList{}
		if info.IsDir() {
			list = &FileList{
				BasePath: storageAbsPath,
				FilePath: filePath,
				FileName: info.Name(),
				FileType: "dir",
				FileSize: "-",
				DateTime: datetimex.FormatDateTime(fileInfo.ModTime()),
			}
			fileList = append(fileList, *list)
			return filepath.SkipDir
		}
		if filePath != storageAbsPath {
			list = &FileList{
				BasePath: storageAbsPath,
				FilePath: filePath,
				FileName: info.Name(),
				FileType: "file",
				FileSize: bytex.FormatFileSize(fileInfo.Size()),
				DateTime: datetimex.FormatDateTime(fileInfo.ModTime()),
			}
			fileList = append(fileList, *list)
		}
		return nil
	})
	if err != nil {
		logx.GetLogger().Sugar().Errorf("Traversal file directory failed!\n%s", err.Error())
		panic(err.Error())
	}
	// 返回目录json数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Get data successfully!",
		"data": fileList,
	})
}
