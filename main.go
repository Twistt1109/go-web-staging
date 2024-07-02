package main

import (
	"fmt"
	"go-web-staging/config"
)

func main() {
	// TODO:
	// 1. 加载配置文件
	if err := config.Init(); err != nil {
		fmt.Println("配置文件加载失败")
		return
	}

	fmt.Println(config.Config)
	// 2. 初始化日志
	// 3. 数据库连接
	// 4. redis连接
	// 5. 路由注册
	// 6. 启动服务
}
