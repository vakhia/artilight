package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/controller/handlers"
	"github.com/vakhia/artilight/internal/middlewares"
	"github.com/vakhia/artilight/pkg/token"
)

func InitCategoryRoutes(router *gin.Engine, jwtServices token.IJwtService, categoryHandler handlers.CategoryHandler) {
	categoryRoutes := router.Group("/api/v1/categories")
	{
		categoryRoutes.GET("", middlewares.Authenticate(jwtServices), categoryHandler.GetAllCategories)
		categoryRoutes.GET("/:id", middlewares.Authenticate(jwtServices), categoryHandler.GetCategoryById)
		categoryRoutes.POST("", middlewares.Authenticate(jwtServices), categoryHandler.CreateCategory)
		categoryRoutes.PUT("/:id", middlewares.Authenticate(jwtServices), categoryHandler.UpdateCategory)
		categoryRoutes.DELETE("/:id", middlewares.Authenticate(jwtServices), categoryHandler.DeleteCategory)
	}
}
