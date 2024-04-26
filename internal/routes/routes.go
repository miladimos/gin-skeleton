package routes

import (
	"gin-skeleton/internal/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// log := logrus.New()
	// r.Use(middle.Logger(log), gin.Recovery())

	// start middlewares

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.NoRoute(middleware.NoRouteHandler())
	router.HandleMethodNotAllowed = true
	router.NoMethod(middleware.NoMethodHandler())

	router.Use(cors.Default())

	// end middlewares

	// routes.AdminRoute(router)

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

			// v1.GET("/posts", controller.PostController.Index)
			// v1.GET("/posts/:id", getArticleOne)
			// v1.POST("/posts", createArticle)
			// v1.PUT("/posts/:id", updateArticle)
			// v1.DELETE("/posts/:id", deleteArticle)

			// authenticated := v1.Use(middleware.AuthMiddleware())
			// {

			// }
		}
	}

	return router
}
