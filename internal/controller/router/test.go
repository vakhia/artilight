package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/middlewares"
	"github.com/vakhia/artilight/pkg/token"
)

func InitTestRoutes(router *gin.Engine, jwtService token.IJwtService) {
	testRoutes := router.Group("/api/v1/test")
	{
		testRoutes.GET("test", middlewares.Authenticate(jwtService), func(context *gin.Context) {
			context.JSON(200, gin.H{"message": "Hello, World!"})
		})
	}
}
