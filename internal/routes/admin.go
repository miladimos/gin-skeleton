package routes

import "github.com/gin-gonic/gin"

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "admin")
		})
	}
}
