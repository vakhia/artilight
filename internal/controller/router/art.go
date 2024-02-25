package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/controller/handlers"
	"github.com/vakhia/artilight/internal/middlewares"
	"github.com/vakhia/artilight/pkg/token"
)

func InitArtRoutes(router *gin.Engine, jwtServices token.IJwtService, artHandler handlers.ArtHandler) {
	categoryRoutes := router.Group("/api/v1/arts")
	{
		categoryRoutes.GET("", middlewares.Authenticate(jwtServices), artHandler.GetAllArts)
		categoryRoutes.GET("/:id", middlewares.Authenticate(jwtServices), artHandler.GetArtById)
		categoryRoutes.POST("", middlewares.Authenticate(jwtServices), artHandler.CreateArt)
		categoryRoutes.PUT("/:id", middlewares.Authenticate(jwtServices), artHandler.UpdateArt)
		categoryRoutes.DELETE("/:id", middlewares.Authenticate(jwtServices), artHandler.DeleteArt)
	}
}
