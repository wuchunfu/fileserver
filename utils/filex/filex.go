package filex

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// IsDir 判断所给路径是否为文件夹
// IsDir returns true if given path is a dir,
// or returns false when it's a directory or does not exist.
func IsDir(filePath string) bool {
	file, err := os.Stat(filePath)
	return err == nil && file.IsDir()
}

// GetFullPath 获取绝对路径
func GetFullPath(path string) string {
	abPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}
	return abPath
}

// IsFile 判断所给路径是否为文件
// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	return !IsDir(filePath)
}

// FilePathExists Judge whether the given path file / folder exists.
func FilePathExists(filePath string) bool {
	_, err := os.Lstat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// Mkdir 创建目录
func Mkdir(filePath string) bool {
	err := os.Mkdir(filePath, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

// MkdirAll 递归创建目录
func MkdirAll(filePath string) bool {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

func IsAbsolutePath(pathStr string) bool {
	goos := runtime.GOOS
	includeColon := strings.Contains(pathStr, string(os.PathListSeparator))
	if includeColon == true || goos == "windows" {
		paths := strings.Replace(pathStr, "\\", "/", -1)
		index := strings.Index(paths, string(os.PathListSeparator))
		return path.IsAbs(paths[index+1:])
	}
	return path.IsAbs(pathStr)
}

// Sha1f return file sha1 encode
func Sha1f(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// ReadFile 读取文件
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer func(fi *os.File) {
		err := fi.Close()
		if err != nil {
			fmt.Println("Close error: ", err)
		}
	}(fi)
	fd, err := io.ReadAll(fi)
	return string(fd)
}

// GetFileSize 获取文件路径
func GetFileSize(filename string) int64 {
	var result int64
	err := filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	if err != nil {
		return 0
	}
	return result
}
