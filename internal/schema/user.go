package schema

import (
	"example/internal/repository"
	"fmt"
	"net/http"
)

// User 映射数据库的结构
type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Account string `json:"account"`
	Role    string `json:"role"`
}

var db = repository.DB

// GetAllUser 获取所有用户
func GetAllUser() Status {
	var users []User
	result := db.Find(users)

	fmt.Printf("%v", result)

	return Status{
		Code:    200,
		Message: "OK",
		Body:    result,
	}
}

// CreateUser 创建用户
func CreateUser(user User) (Status, error) {
	// 创建用户
	result := db.Create(&user)

	if result.Error != nil {
		return Status{
			Code:    http.StatusBadRequest,
			Message: "Error" + result.Error.Error(),
			Body:    result,
		}, result.Error
	}

	return Status{
		Code:    200,
		Message: fmt.Sprintf("Created Row: %d", result.RowsAffected),
		Body:    user.Id,
	}, nil
}
