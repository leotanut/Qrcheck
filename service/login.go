package service

import (
	"webapp-check-in/models"

	"golang.org/x/crypto/bcrypt"
)

/*type LoginService interface {
	LoginUser(email string, password string) bool
}*/

/*func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "bikash.dulal@wesionary.team",
		password: "testing",
	}
}*/

func LoginUser(email string, password string) (models.LoginResp, bool) {

	res := models.LoginResp{}
	models.DB.Table("admins").Select("*").Where("email = ?", email).Scan(&res)
	/*if err != nil {
		return res, false
	}*/
	/*	hashPassInfo, _ := HashPassword(password)
		fmt.Println(hashPassInfo)
		hashPassDB, _ := HashPassword(password)
		fmt.Println(hashPassDB)*/
	/*fmt.Println(password)
	fmt.Println(res.Password)*/
	if CheckPasswordHash(password, res.Password) {
		return res, true
	}
	return res, false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
