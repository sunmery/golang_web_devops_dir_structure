package middleware

import (
	"example/internal/schema"
	"example/pkg/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// AuthRequired 中间件 验证 Token 是否有效
func AuthRequired(c *gin.Context) {
	AuthHead := c.GetHeader("Authorization")
	if AuthHead == "" {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    http.StatusExpectationFailed,
			Message: "请携带Authorization Header,Authorization Header头为空",
			Body:    nil,
		})
		return
	} else if len(AuthHead) <= 7 {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    http.StatusExpectationFailed,
			Message: "请携带正确的JWT的值,Authorization Header值不正确",
			Body:    nil,
		})
		return
	}
	Authorization := c.GetHeader("Authorization") // 获取Authorization值
	token := Authorization[7:]                    // 截取token部分

	jwtKey := []byte(os.Getenv("JWT_KEY"))
	_, verifyErr := helper.ParseToken(token, jwtKey) // 解析JWT
	if verifyErr != nil {
		c.AbortWithStatusJSON(http.StatusOK, schema.Status{
			Code:    http.StatusForbidden,
			Message: "校验token失败, token无效或已过期",
			Body:    nil,
		})
		return
	}

	c.Next()
}
