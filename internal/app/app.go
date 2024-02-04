package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/config"
	"github.com/vakhia/artilight/internal/server"
	"net/http"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin server!"})
	})

	err = server.NewServer(cfg, router).Run()
	if err != nil {
		return
	}
}
