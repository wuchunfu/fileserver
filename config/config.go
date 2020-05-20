package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config/config.yaml"

var InitViper *viper.Viper
var InitConfig Server

func init() {
	vip := viper.New()
	vip.SetConfigFile(defaultConfigFile)
	err := vip.ReadInConfig()
	if err != nil {
		logrus.Errorf("Failed to get config file!")
		panic(err.Error())
	}
	vip.WatchConfig()
	vip.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Config file changed: %s", e.Name)
		if err := vip.Unmarshal(&InitConfig); err != nil {
			logrus.Errorf("Failed to resolve config file!")
			panic(err.Error())
		}
	})
	if err := vip.Unmarshal(&InitConfig); err != nil {
		logrus.Errorf("Failed to resolve config file!")
		panic(err.Error())
	}
	InitViper = vip
}

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
}

type System struct {
	Addr        string `mapstructure:"addr" json:"addr" yaml:"addr"`
	StoragePath string `mapstructure:"storage-path" json:"storagePath" yaml:"storage-path"`
}

type Log struct {
	Path  string `mapstructure:"path" json:"path" yaml:"path"`
	Name  string `mapstructure:"name" json:"name" yaml:"name"`
	Level string `mapstructure:"level" json:"level" yaml:"level"`
}
