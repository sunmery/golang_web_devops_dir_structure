package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	HOST     = os.Getenv("POSTGRES_HOST")     // 主机
	PORT     = os.Getenv("POSTGRES_PORT")     // 端口
	USER     = os.Getenv("POSTGRES_USERNAME") // 用户, 默认为 postgres
	PASSWORD = os.Getenv("POSTGRES_PASSWORD") // 密码
	DATABASE = os.Getenv("POSTGRES_DATABASE") // 数据库
	TimeZone = "Asia/Shanghai"                // 时区
	SslMode  = "disable"                      // 是否加密传输, 默认关闭, enable启用
)

var DB *gorm.DB // pg 客户端
var err error

// InitPostgres 初始化数据库服务
func InitPostgres() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", HOST, USER, PASSWORD, DATABASE, PORT, SslMode, TimeZone)
	DB, err = gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: false,
		}))
	if err != nil {
		panic("启动 Postgres 服务失败!, 错误信息:" + err.Error())
	}
}
