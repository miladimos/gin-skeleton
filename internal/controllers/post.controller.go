package controllers

import (
	"gin-skeleton/internal/models"
	services "gin-skeleton/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostControllerInterface interface {
	Index(context *gin.Context)
	Read(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type PostController struct {
	svc *services.PostService
}

func NewPostController(postService services.PostService) *PostController {
	return &PostController{
		svc: &postService,
	}
}

func (controller PostController) Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"ping": "posts in controller"})
}

func (controller *PostController) Get(context *gin.Context) {

}

func (controller *PostController) Create(context *gin.Context) {
	var post models.Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		panic(err)
	}
	// controller.svc.Create(post)
}
