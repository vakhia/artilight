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

	//Core
	userRepo := repositories.NewUserRepository(db)
	userUseCase := services.NewUserUseCase(userRepo, tokenService)
	userHandler := handlers.NewUserHandler(userUseCase)
	router.InitAuthRoutes(ginRouter, userHandler)
	router.InitTestRoutes(ginRouter, tokenService)

	err = server.NewServer(cfg, ginRouter).Run()
	if err != nil {
		return
	}
}
