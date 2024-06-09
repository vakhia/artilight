package art

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/art/adapters"
	"github.com/vakhia/artilight/internal/art/application"
	"github.com/vakhia/artilight/internal/art/application/command"
	"github.com/vakhia/artilight/internal/art/application/query"
	"github.com/vakhia/artilight/internal/art/domain/repository"
	"github.com/vakhia/artilight/internal/art/ports"
	"github.com/vakhia/artilight/internal/common/config"
	"github.com/vakhia/artilight/internal/common/container"
	"github.com/vakhia/artilight/internal/common/middlewares"
	"gorm.io/gorm"
)

// Module represents the art module and its dependencies.
type Module struct {
	Config     *config.Config
	DB         *gorm.DB
	Router     *gin.Engine
	HttpServer *ports.HttpServer
	Container  *container.Container
}

// NewModule creates a new instance of the art module.
func NewModule(cfg *config.Config, db *gorm.DB, router *gin.Engine, container *container.Container) *Module {
	artRepository := adapters.NewPgSqlArtRepository(db)
	categoryRepository := adapters.NewPgSqlCategoryRepository(db)
	collectionRepository := adapters.NewPgSqlCollectionRepository(db)
	userAdapter := adapters.NewUserAdapter(container.UserRepository)
	app := application.Application{
		Commands: application.Commands{
			CreateArtCommand:        command.NewCreateArtHandler(artRepository, categoryRepository, userAdapter, collectionRepository),
			CreateCategoryCommand:   command.NewCreateCategoryHandler(categoryRepository),
			CreateCollectionCommand: command.NewCollectionHandler(collectionRepository, userAdapter),
		},
		Queries: application.Queries{
			AllArts:       query.NewAllArtsQuery(artRepository),
			GetArt:        query.NewGetArtHandler(artRepository),
			AllCollection: query.NewAllCollectionQuery(collectionRepository),
			AllCategories: query.NewAllCategoriesQuery(categoryRepository),
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

// RegisterRoutes sets up the routes for the art module.
func (m *Module) RegisterRoutes() {
	// Arts routes
	artGroup := m.Router.Group("/api/v1/arts")
	{
		artGroup.GET("", m.HttpServer.GetAllArts)
		artGroup.GET("/:id", m.HttpServer.GetArt)
		artGroup.POST("", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.CreateArt)
	}
	// Categories routes
	categoryGroup := m.Router.Group("/api/v1/categories")
	{
		categoryGroup.POST("", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.CreateCategory)
		categoryGroup.GET("", m.HttpServer.GetAllCategories)
	}
	// Collections routes
	collectionGroup := m.Router.Group("/api/v1/collections")
	{
		collectionGroup.GET("", m.HttpServer.GetAllCollections)
		collectionGroup.POST("", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.CreateCollection)
	}
}

func (m *Module) GetArtRepository() repository.ArtRepository {
	return adapters.NewPgSqlArtRepository(m.DB)
}
