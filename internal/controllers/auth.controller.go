package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miladimos/sanjabi/app/serializers"
)

type AuthenticationController interface {
	Login(context *gin.Context)
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
