package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config *AppConfig

type AppConfig struct {
	Port string `mapstructrue:"port"`
}

func Init() (err error) {
	v := viper.New()
	v.SetConfigFile("config.yml")

	// 读取配置文件
	if err = v.ReadInConfig(); err != nil {
		fmt.Println("load config filed err: ", err)
		return
	}

	// 解析配置文件
	if err = v.Unmarshal(&Config); err != nil {
		fmt.Println("unmarshal config filed err: ", err)
		return
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err = v.Unmarshal(&Config); err != nil {
			fmt.Println("unmarshal config filed err: ", err)
			return
		}
	})
	return
}
