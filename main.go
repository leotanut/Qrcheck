package main

import (
	"net/http"

	controller "webapp-check-in/controllers/loginController"
	userControllers "webapp-check-in/controllers/userController"
	"webapp-check-in/middleware"
	"webapp-check-in/models"
	"webapp-check-in/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	//loginService.LoginUser("csdcsd","csdcsd")
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()
	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	v1 := r.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
	r.GET("/users", userControllers.GetAllUsers)
	r.GET("/dashboard", userControllers.GetUserToday)
	r.POST("/users", userControllers.CreateUser)
	r.Run()
}
