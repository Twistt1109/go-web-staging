package main

import (
	"context"
	"fmt"
	"go-web-staging/internal/app"
	"go-web-staging/pkg/snowflake"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// TODO:
	// 1. 加载配置文件
	if err := app.InitConfig(); err != nil {
		fmt.Println("配置文件加载失败")
		return
	}

	// 2. 初始化日志
	if err := InitLogger(app.Conf.LogConfig); err != nil {
		fmt.Println("日志初始化失败")
		return
	}
	defer zap.L().Sync() // 同步日志

	// 3. 数据库连接
	if err := InitDB(app.Conf.MysqlConfig); err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	defer CloseDB()

	// 4. redis连接
	if err := InitRedis(app.Conf.RedisConfig); err != nil {
		fmt.Println("redis连接失败")
		return
	}
	defer CloseRedis()

	// 5. 生成雪花id
	if err := snowflake.Init(app.Conf.StartTime, app.Conf.MachineID); err != nil {
		fmt.Println("id生成失败")
		return
	}

	// 6. 路由注册
	r := buildRouter()

	// 7. 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		zap.L().Info(fmt.Sprintf("Server is running at %s", srv.Addr) + " ...")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
