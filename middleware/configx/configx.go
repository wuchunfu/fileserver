package configx

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/wuchunfu/fileserver/utils/filex"
	"go.uber.org/zap"
	"os"
)

type System struct {
	Port        string `yaml:"port"`
	StoragePath string `yaml:"storagePath"`
}

// Log log parameters
type Log struct {
	AppName       string `yaml:"appName"`
	Development   bool   `yaml:"development"`
	Level         string `yaml:"level"`
	LogFileDir    string `yaml:"logFileDir"`
	InfoFileName  string `yaml:"infoFileName"`
	WarnFileName  string `yaml:"warnFileName"`
	ErrorFileName string `yaml:"errorFileName"`
	DebugFileName string `yaml:"debugFileName"`
	MaxAge        int    `yaml:"maxAge"`
	MaxBackups    int    `yaml:"maxBackups"`
	MaxSize       int    `yaml:"maxSize"`
}

// YamlSetting global constants are defined and configured by the user according to the file conf.yaml in conf
type YamlSetting struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
}

var (
	Vip        = viper.New()
	ConfigFile = ""
	// ServerSetting global config
	ServerSetting = new(YamlSetting)
	logger        = zap.Logger{}
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if ConfigFile != "" {
		if !filex.FilePathExists(ConfigFile) {
			logger.Sugar().Errorf("No such file or directory: %s", ConfigFile)
			os.Exit(1)
		} else {
			// Use config file from the flag.
			Vip.SetConfigFile(ConfigFile)
			Vip.SetConfigType("yaml")
		}
	} else {
		logger.Sugar().Errorf("Could not find config file: %s", ConfigFile)
		os.Exit(1)
	}
	// If a config file is found, read it in.
	err := Vip.ReadInConfig()
	if err != nil {
		logger.Sugar().Errorf("Failed to get config file: %s", ConfigFile)
	}
	Vip.WatchConfig()
	Vip.OnConfigChange(func(e fsnotify.Event) {
		logger.Sugar().Infof("Config file changed: %s\n", e.Name)
		fmt.Printf("Config file changed: %s\n", e.Name)
		ServerSetting = GetConfig(Vip)
	})
	Vip.AllSettings()
	ServerSetting = GetConfig(Vip)
}

// GetConfig 解析配置文件，反序列化
func GetConfig(vip *viper.Viper) *YamlSetting {
	setting := new(YamlSetting)
	// 解析配置文件，反序列化
	if err := vip.Unmarshal(setting); err != nil {
		logger.Sugar().Errorf("Unmarshal yaml faild: %s", err)
		os.Exit(1)
	}
	return setting
}
