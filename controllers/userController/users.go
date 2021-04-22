package controller

import (
	"net/http"
	"time"
	"webapp-check-in/models"

	"github.com/gin-gonic/gin"
)

type CreateUsersInput struct {
	Name   string `json:"name" binding:"required"`
	Status int    `json:"status" binding:"required"`
}

type DashboardOutput struct {
	Name   string    `json:"name" binding:"required"`
	Status int       `json:"status" binding:"required"`
	Time   time.Time `json:"create_at" binding:"required"`
}

func CreateUser(c *gin.Context) {
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
	var input CreateUsersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.Users{Name: input.Name, Status: input.Status}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetAllUsers(c *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserToday(c *gin.Context) {
	//var users []models.Users
	var dashboard []DashboardOutput
	var currentTime time.Time = time.Now()
	models.DB.Table("users").Select("name", "status", "created_at").Where("DATE(created_at) = ?", currentTime.Format("2006-01-02")).Limit(20).Scan(&dashboard)
	c.JSON(http.StatusOK, gin.H{"data": dashboard})
}

/*func pwhash(pw, salt []byte) ([]byte, error) {
	return scrypt.Key(pw, salt, 32768, 8, 1, 32)
}*/
