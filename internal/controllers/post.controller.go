package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miladimos/sanjabi/app/serializers"
)

type PostController interface {
	Index(context *gin.Context)
	Read(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

func Get(context *gin.Context) {
	var registerRequest serializers.RegisterSerializer
	err := context.ShouldBind(&registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func Index(context *gin.Context) {
	//
}
