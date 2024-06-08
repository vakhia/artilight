package auction

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/auction/adapters"
	"github.com/vakhia/artilight/internal/auction/application"
	"github.com/vakhia/artilight/internal/auction/application/command"
	"github.com/vakhia/artilight/internal/auction/application/query"
	"github.com/vakhia/artilight/internal/auction/ports"
	"github.com/vakhia/artilight/internal/common/config"
	"github.com/vakhia/artilight/internal/common/container"
	"github.com/vakhia/artilight/internal/common/middlewares"
	"gorm.io/gorm"
)

// Module represents the auction module and its dependencies.
type Module struct {
	Config     *config.Config
	DB         *gorm.DB
	Router     *gin.Engine
	HttpServer *ports.HttpServer
	Container  *container.Container
}

// NewModule creates a new instance of the auction module.
func NewModule(cfg *config.Config, db *gorm.DB, router *gin.Engine, container *container.Container) *Module {
	auctionRepository := adapters.NewPgSqlAuctionRepository(db)
	bidRepository := adapters.NewPgSqlBidRepository(db)
	userService := adapters.NewUserAdapter(container.UserRepository)
	itemService := adapters.NewArtAdapter(container.ArtRepository)
	app := application.Application{
		Commands: application.Commands{
			CreateAuctionCommand: command.NewCreateAuctionCommandHandler(auctionRepository, userService, itemService),
			UpdateAuctionCommand: command.NewUpdateAuctionCommandHandler(auctionRepository),
			CreateBidCommand:     command.NewCreateBidCommandHandler(bidRepository, userService, itemService, auctionRepository),
		},
		Queries: application.Queries{
			GetAuctionQuery: query.NewGetAuctionHandler(auctionRepository),
		},
	}

	httpServer := ports.NewHttpServer(app)

	return &Module{
		Config:     cfg,
		DB:         db,
		Router:     router,
		HttpServer: httpServer,
		Container:  container,
	}
}

// RegisterRoutes sets up the routes for the auction module.
func (m *Module) RegisterRoutes() {
	// Auction routes
	auctionGroup := m.Router.Group("/api/v1/auctions")
	{
		auctionGroup.GET("/:id", m.HttpServer.GetAuction)
		auctionGroup.PUT("/:id", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.UpdateAuction)
		auctionGroup.POST("", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.CreateAuction)
	}

	bidGroup := m.Router.Group("/api/v1/bids")
	{
		bidGroup.POST("", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.CreateBid)
	}
}
