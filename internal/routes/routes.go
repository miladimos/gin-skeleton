package routes

import (
	"gin-skeleton/internal/controllers"
	"gin-skeleton/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// if !config.App.Debug {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	router := gin.Default()

	// Start logger
	// log := logrus.New()
	// log.SetReportCaller(true)
	// logFile, _ := os.Create("log.log")
	// multiLogWriter := io.MultiWriter(logFile, os.Stdout)
	// log.SetOutput(multiLogWriter)
	// r.Use(middle.Logger(log))
	// End logger

	// start middlewares

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.NoRoute(middleware.NoRouteHandler())
	router.HandleMethodNotAllowed = true
	router.NoMethod(middleware.NoMethodHandler())
	// router.Use(cors.Default())

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

				// ctx.POST("/emails/verify",authController.Register)

				// ctx.POST("/passwords/forgot",authController.Register)
				// ctx.GET("/passwords/reest",authController.Register)
			})

			v1.GET("/tags", controllers.TagList)
			v1.GET("/tags/:id", controllers.TagShow)
			v1.POST("/tags", controllers.TagCreate)
			v1.PUT("/posts/:id", controllers.TagUpdate)
			v1.DELETE("/posts/:id", controllers.TagDelete)

			// authenticated := v1.Use(middleware.AuthMiddleware())
			// {
			// authenticated.POST("/accounts", getArticleOne)

			// }
		}
	}

	return router
}
