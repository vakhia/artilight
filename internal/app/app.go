package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/config"
	"github.com/vakhia/artilight/internal/controller/handlers"
	"github.com/vakhia/artilight/internal/controller/router"
	"github.com/vakhia/artilight/internal/database"
	"github.com/vakhia/artilight/internal/repositories"
	"github.com/vakhia/artilight/internal/server"
	"github.com/vakhia/artilight/internal/services"
	"github.com/vakhia/artilight/pkg/token"
	"net/http"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	ginRouter := gin.Default()

	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin server!"})
	})
	db := database.Open(cfg)
	defer database.Close(db)

	//Services
	tokenService := token.NewJWTService(cfg)

	//User
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, tokenService)
	userHandler := handlers.NewUserHandler(userService)

	//Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	//Art
	artRepo := repositories.NewArtRepository(db)
	artService := services.NewArtService(artRepo)
	artHandler := handlers.NewArtHandler(artService)

	//Routes
	router.InitAuthRoutes(ginRouter, userHandler)
	router.InitTestRoutes(ginRouter, tokenService)
	router.InitCategoryRoutes(ginRouter, tokenService, categoryHandler)
	router.InitArtRoutes(ginRouter, tokenService, artHandler)

	err = server.NewServer(cfg, ginRouter).Run()
	if err != nil {
		return
	}
}
