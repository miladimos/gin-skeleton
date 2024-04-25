package routes

import (
	"gin-skeleton/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// setup middlewares
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
				// ctx.Get("/register",authController.Register)
				// ctx.Get("/login",authController.Register)
				// ctx.Get("/logout",authController.Register)
			})

			// authenticated := v1.Use(middleware.AuthMiddleware())
			// {

			// }
		}
	}

	return router
}
