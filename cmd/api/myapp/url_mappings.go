package app

import user "crud_with_TG/Golang-Poc-TG/internal/social/controllers"

func mapUrls() {
	router.POST("/users", user.Create)
	router.PATCH("/users/:id", user.UpdateUser)
	router.DELETE("/users/:id", user.DeleteUser)
	router.GET("/users/:id", user.GetUser)
	router.GET("/users", user.GetAllUser)
	router.GET("/friendship", user.GetAllFriendship)
	router.GET("/ping", user.Hello)
}
