package main

import (
	"fmt"
	"go-web-staging/logger"
	"go-web-staging/setting"
)

func main() {
	// TODO:
	// 1. 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Println("配置文件加载失败")
		return
	}

	// 2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Println("日志初始化失败")
		return
	}
	// 3. 数据库连接
	// 4. redis连接
	// 5. 路由注册
	// 6. 启动服务
}
