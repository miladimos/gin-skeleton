package controllers

import (
	"gin-skeleton/internal/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
	Logout(context *gin.Context)
}

func Register(context *gin.Context) {
	var registerRequest serializers.RegisterSerializer
	err := context.ShouldBind(&registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func Login(context *gin.Context) {
	//
}

func Logout(context *gin.Context) {
	//
}
