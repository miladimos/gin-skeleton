package controllers

import "github.com/gin-gonic/gin"

func ForgotPassword(c *gin.Context) {
	var email string = c.Param("email")

}
