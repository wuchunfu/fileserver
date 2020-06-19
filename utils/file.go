package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 创建目录
func Mkdir(filePath string) bool {
	err := os.Mkdir(filePath, os.ModePerm)
	if err != nil {
		logrus.Errorf("mkdir %s failed. %s", filePath, err)
		return false
	}
	return true
}

// 递归创建目录
func MkdirAll(filePath string) bool {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		logrus.Errorf("mkdir %s failed. %s", filePath, err)
		return false
	}
	return true
}

// 判断所给路径文件/文件夹是否存在
func IsExistPath(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
