package handler

import (
	"example/internal/schema"
	"example/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

type RoleInter interface {
	Auth(c *gin.Context)
}

type Role struct{}

// Token 给用户返回 Token
func (r Role) Token(c *gin.Context) {
	// JWT结构体赋值
	claim := schema.Claim{
		Name:    "test user", // 用户名
		Account: "12345",     // 用户ID
		Role:    "admin",     // 用户权限
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	jwtKey := []byte(os.Getenv("SIGNING_KEY")) // 读取配置文件获取jwtKey私钥
	tokenString, err := helper.GenerateToken(jwt.SigningMethodES256, claim, jwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.Header("token", tokenString)
}
