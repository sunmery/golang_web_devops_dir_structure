package router

import (
	"example/docs"
	"example/internal/handler"
	"example/internal/middleware"
	"example/internal/schema"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path/filepath"
)

func WebServer(host string, port string) {
	// 配置 Web 服务
	r := gin.Default()

	// CORS 跨域处理配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"} // 允许的请求方法
	corsConfig.AllowAllOrigins = true                                           // 是否允许所有IP的请求

	//测试服务
	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, schema.Status{
			Code:    http.StatusOK,
			Message: "OK",
			Body:    fmt.Sprintf("Hello! Welcome Test Gin Web Server API! You IP is %s", c.ClientIP()),
		})
	})

	// http://127.0.0.1:<prot>/swagger/index.html
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1", middleware.AuthRequired)
	{
		user := v1.Group("/user")
		{
			user.GET("", handler.User{}.GetUser)
			user.PUT("/create", handler.User{}.CreateUser)
		}
	}

	// Swagger UI 文档服务器
	// https://github.com/swaggo/gin-swagger
	// http://127.0.0.1:<port>/swagger/doc.json
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", 4001)), // 指向 API 定义的网址
		ginSwagger.DocExpansion("list"),         // 控制操作和标记的默认展开设置。它可以是 `list`（仅展开标签）、`full`（展开标签和操作）或 `none`（不展开任何内容）
		ginSwagger.DeepLinking(true),            // 如果设置为 true，则为标签和操作启用深层链接。有关更多信息，请参阅深层链接文档
		ginSwagger.DefaultModelsExpandDepth(-1), // 模型的默认扩展深度（设置为 -1 将完全隐藏模型)
		ginSwagger.InstanceName("swagger"),      // swagger文档的实例名称。如果要在一个 gin 路由器上部署多个不同的 swagger 实例，请确保每个实例都有一个唯一的名称（使用 --instanceName 参数生成带有 swag init 的 swagger 文档）。
		ginSwagger.PersistAuthorization(false),  // 如果设置为 true，它将保留授权数据，并且在浏览器关闭/刷新时不会丢失
		ginSwagger.Oauth2DefaultClientID(""),    // 如果设置，它将用于预填充 OAuth2 授权对话框的client_id字段
	))

	// HTTPS SSL 配置
	if port == "443" {
		sslCertPath := filepath.Join("internal/config", "ssl", "server.crt")
		sslKeyPath := filepath.Join("internal/config", "ssl", "server.key")

		slog.Infof("启动https://%s:%s服务", host, port)
		if err := r.RunTLS(fmt.Sprintf("0.0.0.0:%s", port), sslCertPath, sslKeyPath); err != nil {
			slog.Errorf("运行gin服务失败,请检查端口%s是否被占用,错误信息:%s", port, err.Error())
		}
	}

	// HTTP 开发模式
	slog.Info(fmt.Sprintf("启动http://%s:%s服务", host, port))
	if err := r.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		slog.Errorf("运行gin服务失败,请检查端口%s是否被占用,错误信息:%s", port, err.Error())
	}
}
