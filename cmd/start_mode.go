package cmd

import (
	"example/internal/router"
	"flag"
)

var MODE string
var PORT string
var HOST string
var PRODUCTION = []string{"pro", "prod", "production"}
var DEVELOPMENT = []string{"dev", "develop", "development"}

// StartMode 程序启动模式
func StartMode() {
	flag.StringVar(&MODE, "mode", "dev", "环境")
	flag.StringVar(&PORT, "port", "4000", "端口")
	flag.StringVar(&HOST, "host", "0.0.0.0", "主机")
	flag.Parse()
}

// InitServer 程序启动服务列表
func InitServer() {
	//repository.InitPostgres()    // 数据库
	router.WebServer(HOST, PORT) // Web 服务
}
