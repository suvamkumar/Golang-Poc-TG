package user

import (
	users "crud_with_TG/Golang-Poc-TG/internal/social/domain"
	"crud_with_TG/Golang-Poc-TG/internal/social/services"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"
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

//CreateManyUser ...
func CreateManyUser(c *gin.Context) {
	var user []users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
	}
	result, err := services.UserService.CreateManyUser(user)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//SyncWithDB ...
func SyncWithDB(c *gin.Context) {
	result, err := services.UserService.SyncDBWithTG()
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
