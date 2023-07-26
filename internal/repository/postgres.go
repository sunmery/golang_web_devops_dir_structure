package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB // pg 客户端
var PostgresErr error

// InitPostgresConnect 初始化数据库服务
func InitPostgresConnect() {
	HOST := os.Getenv("POSTGRES_HOST")          // 主机
	USER := os.Getenv("POSTGRES_USER")          // 用户, 默认为 postgres
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")  // 密码
	PORT := os.Getenv("POSTGRES_PORT")          // 端口
	DATABASE := os.Getenv("POSTGRES_DATABASE")  // 数据库
	SslMode := os.Getenv("POSTGRES_SSL_MODE")   // 时区
	TimeZone := os.Getenv("POSTGRES_TIME_ZONE") // 是否加密传输, 默认关闭, enable启用

	// 连接数据库
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		HOST, USER, PASSWORD, DATABASE, PORT, SslMode, TimeZone,
	)
	PostgresDB, PostgresErr = gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: false,
		}))
	if PostgresErr != nil {
		panic("启动 Postgres 服务失败!, 错误信息:" + PostgresErr.Error())
	}
}
