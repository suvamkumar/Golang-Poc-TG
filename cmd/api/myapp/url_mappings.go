package app

import user "crud_with_gin_gonic/internal/users/controllers"

func mapUrls() {
	router.POST("/users", user.Create)
	router.PATCH("/users/:id", user.UpdateUser)
	router.DELETE("/users/:id", user.DeleteUser)
	router.GET("/users/:id", user.GetUser)
	router.GET("/users", user.GetAllUser)
	router.GET("/ping", user.Hello)
}
