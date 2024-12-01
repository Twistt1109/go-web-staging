package app

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Conf *AppConfig

type AppConfig struct {
	Name         string `mapstructure:"name"`
	ErrorFile    string `mapstructure:"error_file"`
	Version      string `mapstructure:"version"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	Secret       string `mapstructure:"secret"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LogConfig struct {
	Console    bool   `mapstructure:"console"`
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func InitConfig() (err error) {
	// 加载 .env 文件中的配置到环境变量
	// 如果是嵌套的结构体, 文件中参数命名格式: 以.连接(如MYSQL.HOST)
	godotenv.Load("../.env")

	v := viper.New()
	v.SetConfigFile("../configs/app.yml") // 指定配置文件路径

	// v.SetDefault("error_file", "../config/errors.yaml")

	v.SetEnvPrefix("app") // 设置环境变量的前缀, 加上前缀, 且大写
	v.AutomaticEnv()      // 自动读取环境变量

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

	return
}
