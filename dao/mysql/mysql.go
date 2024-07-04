package mysql

import (
	"fmt"
	"go-web-staging/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *setting.MysqlConfig) (err error) {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("连接数据库失败", zap.Error(err))
		return err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}

// 关闭数据库连接
func Close() {
	db.Close()
}
