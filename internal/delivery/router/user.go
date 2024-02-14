package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/delivery/handlers"
)

func InitAuthRoutes(router *gin.Engine, userHandler handlers.UserHandler) {
	userRoutes := router.Group("/api/v1/auth")
	{
		userRoutes.POST("register", userHandler.Register)
		userRoutes.POST("login", userHandler.Login)
	}
}
