package controller

import "github.com/gin-gonic/gin"

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}

}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {}

}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func HashPassword(password string) string {
	return ""
}

func verifyPassword(userPassword string, password string) (bool, string) {
	return false, ""
}
