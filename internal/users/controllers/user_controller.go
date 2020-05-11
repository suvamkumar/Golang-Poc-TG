package user

import (
	users "crud_with_gin_gonic/internal/users/domain"
	"crud_with_gin_gonic/internal/users/services"
	"crud_with_gin_gonic/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Create ...
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
	}
	result, err := services.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//GetUser ...
func GetUser(c *gin.Context) {
	id := c.Param("id")
	result, err := services.UserService.GetUser(id)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//GetAllUser ...
func GetAllUser(c *gin.Context) {
	result, err := services.UserService.GetAllUser()
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//UpdateUser ...
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
	}
	result, err := services.UserService.UpdateUser(user, id)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//DeleteUser ...
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.UserService.DeleteUser(id); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, map[string]string{id: "deleted successfully"})
}

//Hello ...
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
