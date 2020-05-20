package common

import "fileserver/config"

// 初始化配置文件
var initConfig = config.InitConfig

// 访问地址
var Addr = initConfig.System.Addr

// 文件上传后的存储路径
var StoragePath = initConfig.System.StoragePath

// 日志级别
var LogLevel = initConfig.Log.Level

// 文件路径
var LogPath = initConfig.Log.Path

// 文件名
var LogName = initConfig.Log.Name
