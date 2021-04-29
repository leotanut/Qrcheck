package main

import (
	"net/http"

	adminControllers "webapp-check-in/controllers/adminController"
	controller "webapp-check-in/controllers/loginController"
	userControllers "webapp-check-in/controllers/userController"
	"webapp-check-in/middleware"
	"webapp-check-in/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//loginService.LoginUser("csdcsd","csdcsd")

	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()
	r.POST("/login", controller.Login)

	v1 := r.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
		v1.GET("/users", userControllers.GetAllUsers)
	}
	r.POST("/create_admin", adminControllers.CreateAdmin)
	//r.GET("/users", userControllers.GetAllUsers)
	r.GET("/dashboard", userControllers.GetUserToday)
	//r.POST("/users", userControllers.CreateUser)
	r.Run(":8080")
}
