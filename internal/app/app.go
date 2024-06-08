package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/art"
	"github.com/vakhia/artilight/internal/auction"
	"github.com/vakhia/artilight/internal/common/config"
	"github.com/vakhia/artilight/internal/common/container"
	"github.com/vakhia/artilight/internal/common/database"
	"github.com/vakhia/artilight/internal/common/fileuploader"
	"github.com/vakhia/artilight/internal/common/ws"
	"github.com/vakhia/artilight/internal/user"
	"github.com/vakhia/artilight/pkg/token"
	"net/http"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	ginRouter := gin.Default()

	// Set up CORS middleware
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
		AllowWebSockets:  true,
	}))

	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin server!"})
	})
	db := database.Open(cfg)
	defer database.Close(db)

	// Initialize the Google Cloud Storage
	uploader, err := fileuploader.NewGCSUploader(cfg.GCS.Bucket)
	if err != nil {
		panic(err)
	}

	// Initialize the JwtService/
	jwtService := token.NewJWTService(cfg)

	// Initialize the DI container.
	container := container.NewContainer()
	container.JwtService = jwtService
	container.Storage = uploader

	// Initialize and register the user module.
	userModule := user.NewModule(cfg, db, ginRouter, container)
	userModule.RegisterRoutes()
	container.UserRepository = userModule.GetUserRepository()

	// Initialize and register the art module.
	artModule := art.NewModule(cfg, db, ginRouter, container)
	artModule.RegisterRoutes()
	container.ArtRepository = artModule.GetArtRepository()

	auctionModule := auction.NewModule(cfg, db, ginRouter, container)
	auctionModule.RegisterRoutes()

	// WebSocket route
	ginRouter.GET("/ws", ws.HandleWebSocket)

	// Start handling messages
	go ws.BroadcastToClients()

	// Start the server.
	if err := ginRouter.Run(); err != nil {
		panic(err)
	}
}
