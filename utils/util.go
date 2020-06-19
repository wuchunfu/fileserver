package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
)

// 检查是否配置环境变量
func CheckFsHome() (string, bool) {
	installPath := os.Getenv("FS_HOME")
	if installPath == "" {
		logrus.Warn("Please add FS_HOME to environment variable, FS_HOME is the installation directory.")
		os.Exit(-1)
		return "", false
	}
	return installPath, true
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
