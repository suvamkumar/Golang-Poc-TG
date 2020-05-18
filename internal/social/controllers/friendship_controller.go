package user

import (
	users "crud_with_TG/Golang-Poc-TG/internal/social/domain"
	"crud_with_TG/Golang-Poc-TG/internal/social/services"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateFriend ...
func CreateFriend(c *gin.Context) {
	var friend users.Friendship
	if err := c.ShouldBindJSON(&friend); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
	}
	result, err := services.FriendshipService.CreateUser(friend)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//CreateManyFriends ...
func CreateManyFriends(c *gin.Context) {
	var friend []users.Friendship
	if err := c.ShouldBindJSON(&friend); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
	}
	result, err := services.FriendshipService.CreateManyEdges(friend)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//GetFriends ...
func GetFriends(c *gin.Context) {
	id := c.Param("id")
	result, err := services.UserService.GetUser(id)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//GetAllFriendship ...
func GetAllFriendship(c *gin.Context) {
	result, err := services.FriendshipService.GetAllUser()
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

//UpdateFriends ...
func UpdateFriends(c *gin.Context) {
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

//DeleteFriends ...
func DeleteFriends(c *gin.Context) {
	id := c.Param("id")
	if err := services.UserService.DeleteUser(id); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, map[string]string{id: "deleted successfully"})
}
