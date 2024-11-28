package main

import (
	"fmt"
	"go-web-staging/internal/app"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(cfg *app.MysqlConfig) (err error) {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	fmt.Println(dsn)
	// db, err = sqlx.Connect("mysql", dsn)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		zap.L().Error("连接数据库失败", zap.Error(err))
		return err
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}

// 关闭数据库连接
func CloseDB() {
	db, _ := db.DB()
	db.Close()
}
