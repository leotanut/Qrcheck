package middleware

import (
	"fmt"
	"net/http"
	"webapp-check-in/service"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		/*if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} */
		//	if len(tokenString) == 0 { fmt.Println("kkkk") }
		//fmt.Println(tokenString)
		if !token.Valid {
			fmt.Println("testing")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
