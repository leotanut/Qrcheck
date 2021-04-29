package controller

import (
	"fmt"
	"net/http"
	"webapp-check-in/models"
	"webapp-check-in/service"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

func Login(ctx *gin.Context) {
	var jwtService service.JWTService = service.JWTAuthService()
	var credential models.LoginRequest
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(credential)
	/*err := ctx.ShouldBind(&credential)
	if err != nil {

	}*/

	data, isUserAuthenticated := service.LoginUser(credential.Email, credential.Password)
	fmt.Println(data)
	//fmt.Println(isUserAuthenticated)
	if isUserAuthenticated {
		ctx.JSON(http.StatusOK, gin.H{"token": jwtService.GenerateToken(string(data.ID), true)})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Wrong username or password "})
	}

}
