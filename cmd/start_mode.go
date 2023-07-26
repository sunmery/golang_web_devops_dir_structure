package cmd

import (
	"example/internal/repository"
	"example/internal/router"
	"flag"
)

var (
	MODE        string
	PORT        string
	HOST        string
	FILE        string
	PRODUCTION  = []string{"pro", "prod", "production"}
	DEVELOPMENT = []string{"dev", "develop", "development"}
)

// StartMode 程序启动模式
func StartMode() {
	flag.StringVar(&MODE, "mode", "dev", "环境")
	flag.StringVar(&PORT, "port", "4000", "端口")
	flag.StringVar(&HOST, "host", "0.0.0.0", "主机")
	flag.StringVar(&FILE, "file", "true", "使用配置文件覆盖默认配置")
	flag.Parse()
}

// InitServer 程序启动服务列表
func InitServer() {
	//repository.InitPostgresConnect()    // Postgres数据库
	repository.InitMysqlConnect() // Mysql数据库
	router.WebServer(HOST, PORT)  // Web 服务
}
