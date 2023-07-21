package schema

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claim JWT 结构体
type Claim struct {
	Name    string `json:"name"`
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}
