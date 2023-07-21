package handler

import (
	"example/internal/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInter interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type User struct{}

//	@BasePath	/api/v1

// GetUser PingExample godoc
//
//	@Summary	获取用户
//	@Schemes
//	@Description	do ping
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/api/v1/user [get]
func (u User) GetUser(c *gin.Context) {
	query := c.Query("user")
	var result any
	if query == "all" {
		result = schema.GetAllUser()
		return
	}
	c.JSON(http.StatusOK, result)
}

//	@BasePath	/api/v1

// CreateUser PingExample godoc
//
//	@Summary	创建用户
//	@Schemes
//	@Description	Create User
//	@Param			ID		path	int		true	"id"
//	@Param			Name	path	string	true	"name"
//	@Param			Account	path	string	true	"account"
//	@Param			Role	path	string	true	"role"
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/api/v1/user/create [put]
func (u User) CreateUser(c *gin.Context) {
	var user schema.User
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, schema.Status{
			Body:    "",
			Message: "",
			Code:    http.StatusBadRequest,
		})
		return
	}

	result, createdErr := schema.CreateUser(user)
	if createdErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, schema.Status{
			Body:    "",
			Message: "",
			Code:    http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
