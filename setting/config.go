package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf *AppConfig

type AppConfig struct {
	Port         string `mapstructure:"port"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
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
	if err = v.Unmarshal(&Conf); err != nil {
		fmt.Println("unmarshal config filed err: ", err)
		return
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err = v.Unmarshal(&Conf); err != nil {
			fmt.Println("unmarshal config filed err: ", err)
			return
		}
	})
	fmt.Println(Conf.LogConfig)

	return
}
