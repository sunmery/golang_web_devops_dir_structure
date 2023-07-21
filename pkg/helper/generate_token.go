package helper

import (
	"example/internal/schema"
	"github.com/golang-jwt/jwt/v5"
)

/** 加密用户字段
 * @description 通过JWT进行加密用户的某些字段
 * @since 2023/1/215:55
 * @param method 加密方式
 * @param claim JWT结构体
 * @param jwtKey JWT私钥
 * @return 完整JWT字符串或者生成JWT的错误信息
 *  */

// GenerateToken 生成 Token 字符串
func GenerateToken(method jwt.SigningMethod, claim schema.Claim, jwtKey []byte) (string, error) {
	// NewWithClaims: 生产header的类型与算法和可解码的body体
	// SignedString 签名, 构成完成的JWT并返回给token变量
	token, generateErr := jwt.NewWithClaims(method, claim).SignedString(jwtKey)
	if generateErr != nil {
		panic(generateErr)
		return "", nil
	}

	return token, nil
}
