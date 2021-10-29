package config

import (
	"fileserver/middleware/configx"
	"fileserver/run"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:          "config",
	SilenceUsage: true,
	Short:        "Get Application config info",
	Example:      "FileServer config -f conf/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		run.Run()
	},
}

func init() {
	cobra.OnInitialize(configx.InitConfig)

	StartCmd.PersistentFlags().StringVarP(&configx.ConfigFile, "configFile", "f", "conf/config.yaml", "config file")
	// 必须配置项
	_ = StartCmd.MarkFlagRequired("configFile")

	// 使用viper可以绑定flag

	// 设置默认值
}
