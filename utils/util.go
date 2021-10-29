package utils

import (
	"github.com/wuchunfu/fileserver/middleware/logx"
	"os"
)

// CheckFsHome 检查是否配置环境变量
func CheckFsHome() (string, bool) {
	installPath := os.Getenv("FS_HOME")
	if installPath == "" {
		logx.GetLogger().Sugar().Warn("Please add FS_HOME to environment variable, FS_HOME is the installation directory.")
		os.Exit(-1)
		return "", false
	}
	return installPath, true
}
