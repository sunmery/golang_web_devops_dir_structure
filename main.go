package main

import (
	"example/cmd"
	"example/pkg/helper"
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
	"os"
)

// 初始化
func init() {
	// dev 开发模式， 默认为 :4000 端口
	// production 默认为 :443 端口
	cmd.StartMode() // go程序启动模式
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:4001
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
//func main() {
//	// 可在 Dockerfile 把新的start.yml文件覆盖旧文件方式启动程序
//	if cmd.FILE == "true" {
//		err := godotenv.Load("./config/start.yml")
//		if err != nil {
//			slog.Error(err)
//		}
//		cmd.MODE = os.Getenv("MODE")
//		cmd.PORT = os.Getenv("PORT")
//		cmd.HOST = os.Getenv("HOST")
//	}
//	// 开发模式, 不输入任何参数默认启动的模式
//	if helper.IsInSlice(cmd.MODE, cmd.DEVELOPMENT) {
//		slog.Info("读取开发环境配置文件")
//
//		err := godotenv.Load("./pkg/config/db.development.yaml")
//		if err != nil {
//			slog.Error(err)
//		}
//
//		// 初始化数据库, Web服务
//		cmd.InitServer()
//	} else if helper.IsInSlice(cmd.MODE, cmd.PRODUCTION) {
//		slog.Info("读取生产环境配置文件")
//
//		err := godotenv.Load("./pkg/config/db.production.yaml")
//		if err != nil {
//			slog.Error(err)
//		}
//
//		cmd.InitServer()
//	}
//
//	// 传递错误的参数会报错
//	slog.Error("请传递正确的参数")
//	os.Exit(2)
//}

func main() {
	// 可在 Dockerfile 把新的start.yml文件覆盖旧文件方式启动程序
	if cmd.FILE == "true" {
		err := godotenv.Load("./config/start.yml")
		if err != nil {
			slog.Error(err)
		}
		cmd.MODE = os.Getenv("MODE")
		cmd.PORT = os.Getenv("PORT")
		cmd.HOST = os.Getenv("HOST")
	}

	if helper.IsInSlice(cmd.MODE, cmd.PRODUCTION) {
		slog.Info("读取生产环境配置文件")

		// 读取生产模式的配置文件
		err := godotenv.Load("./pkg/config/db.production.yaml")
		if err != nil {
			slog.Error(err)
		}
		cmd.InitServer()
		return
	}

	slog.Info("读取开发环境配置文件")

	// 开发模式, 不输入任何参数默认启动的模式
	err := godotenv.Load("./pkg/config/db.development.yaml")
	if err != nil {
		slog.Error(err)
	}

	// 初始化数据库, Web服务
	cmd.InitServer()

	// 传递错误的参数会报错
	slog.Error("请传递正确的参数")
	os.Exit(2)
}
