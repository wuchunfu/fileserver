package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wuchunfu/fileserver/middleware/configx"
	"github.com/wuchunfu/fileserver/middleware/logx"
	"github.com/wuchunfu/fileserver/utils/bytex"
	"github.com/wuchunfu/fileserver/utils/datetimex"
	"github.com/wuchunfu/fileserver/utils/filetypex"
	"github.com/wuchunfu/fileserver/utils/filex"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

type FileList struct {
	BasePath   string `json:"basePath"`
	ParentPath string `json:"parentPath"`
	FilePath   string `json:"filePath"`
	FileName   string `json:"fileName"`
	IsFile     bool   `json:"isFile"`
	FileType   string `json:"fileType"`
	FileSize   string `json:"fileSize"`
	SuffixName string `json:"suffixName"`
	DateTime   string `json:"dateTime"`
}

// List 文件列表
func List(ctx *gin.Context) {
	setting := configx.ServerSetting
	storagePath := setting.System.StoragePath
	storageAbsPath, _ := filepath.Abs(storagePath)
	isExistPath := filex.FilePathExists(storageAbsPath)
	if !isExistPath {
		filex.MkdirAll(storageAbsPath)
	}

	fileList := &[]FileList{}
	parentPath := ctx.Query("parentPath")
	logx.GetLogger().Sugar().Infof("parentPath: %s", parentPath)
	if parentPath == "" {
		fileList = ListFolder(storageAbsPath)
	} else {
		fileList = ListFolder(parentPath)
	}
	// 返回目录json数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Get data successfully!",
		"data": fileList,
	})
}

func ChangeFolder(ctx *gin.Context) {
	fileList := &[]FileList{}

	parentPath := ctx.Query("parentPath")
	logx.GetLogger().Sugar().Infof("parentPath: %s", parentPath)
	if parentPath != "" {
		parentPath := filepath.Dir(parentPath)
		fileList = ListFolder(parentPath)
	}
	// 返回目录json数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Get data successfully!",
		"data": fileList,
	})
}

// ListFolder 列出文件夹中的文件夹及文件
func ListFolder(storageAbsPath string) *[]FileList {
	setting := configx.ServerSetting
	storagePath := setting.System.StoragePath
	basePath, _ := filepath.Abs(storagePath)

	fileList := make([]FileList, 0)
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
			fileName := info.Name()
			dateTime := datetimex.FormatDateTime(fileInfo.ModTime())
			list = &FileList{
				BasePath:   basePath,
				ParentPath: storageAbsPath,
				FilePath:   filePath,
				FileName:   fileName,
				IsFile:     false,
				FileSize:   "-",
				DateTime:   dateTime,
			}
			fileList = append(fileList, *list)
			return filepath.SkipDir
		}
		if filePath != storageAbsPath {
			fileName := info.Name()
			suffixName := strings.ToLower(path.Ext(fileName))
			fileType := filetypex.FileType(suffixName)
			fileSize := bytex.FormatFileSize(fileInfo.Size())
			dateTime := datetimex.FormatDateTime(fileInfo.ModTime())
			list = &FileList{
				BasePath:   basePath,
				ParentPath: storageAbsPath,
				FilePath:   filePath,
				FileName:   fileName,
				IsFile:     true,
				FileType:   fileType,
				SuffixName: suffixName,
				FileSize:   fileSize,
				DateTime:   dateTime,
			}
			fileList = append(fileList, *list)
		}
		return nil
	})
	if err != nil {
		logx.GetLogger().Sugar().Errorf("Traversal file directory failed!\n%s", err.Error())
		panic(err.Error())
	}
	return &fileList
}
