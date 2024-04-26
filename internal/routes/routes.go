package routes

import (
	"gin-skeleton/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// log := logrus.New()
	// r.Use(middle.Logger(log), gin.Recovery())

	// setup middlewares

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.NoRoute(middleware.NoRouteHandler())
	router.HandleMethodNotAllowed = true
	router.NoMethod(middleware.NoMethodHandler())

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/ping", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
			})

			v1.Group("/auth", func(ctx *gin.Context) {
				// ctx.POST("/register",authController.Register)
				// ctx.POST("/login",authController.Register)
				// ctx.POST("/logout",authController.Register)
			})

			// router.GET("/posts", getArticles)
			// router.GET("/posts/:id", getArticleOne)
			// router.POST("/posts", createArticle)
			// router.PUT("/posts/:id", updateArticle)
			// router.DELETE("/posts/:id", deleteArticle)

			// authenticated := v1.Use(middleware.AuthMiddleware())
			// {

			// }
		}
	}

	return router
}
