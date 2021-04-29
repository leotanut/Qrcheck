package controller

import (
	"net/http"
	"webapp-check-in/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateAdminInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/*type DashboardOutput struct {
	Name   string    `json:"name" binding:"required"`
	Status int       `json:"status" binding:"required"`
	Time   time.Time `json:"create_at" binding:"required"`
}*/

func CreateAdmin(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	/*func handler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if (r.Method == "OPTIONS") {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
		} else {
			// Your code goes here
		}
	}*/
	// Validate input
	var input CreateAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPass, _ := HashPassword(input.Password)
	// Create user
	user := models.Admins{Email: input.Email, Password: hashPass}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetAllUsers(c *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*func pwhash(pw, salt []byte) ([]byte, error) {
	return scrypt.Key(pw, salt, 32768, 8, 1, 32)
}*/
