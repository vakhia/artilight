package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/common/config"
	"github.com/vakhia/artilight/internal/common/container"
	"github.com/vakhia/artilight/internal/common/middlewares"
	"github.com/vakhia/artilight/internal/user/adapters"
	"github.com/vakhia/artilight/internal/user/application"
	"github.com/vakhia/artilight/internal/user/application/command"
	"github.com/vakhia/artilight/internal/user/application/query"
	"github.com/vakhia/artilight/internal/user/domain/repository"
	"github.com/vakhia/artilight/internal/user/ports"
	"github.com/vakhia/artilight/pkg/token"
	"gorm.io/gorm"
)

type Module struct {
	Config     *config.Config
	DB         *gorm.DB
	Router     *gin.Engine
	HttpServer *ports.HttpServer
	Container  *container.Container
}

func NewModule(cfg *config.Config, db *gorm.DB, router *gin.Engine, container *container.Container) *Module {
	userRepository := adapters.NewPgSqlUserRepository(db)
	tokenService := token.NewJWTService(cfg)
	app := application.Application{
		Commands: application.Commands{
			CreateUser:   command.NewCreateUserHandler(userRepository),
			LoginUser:    command.NewLoginUserHandler(userRepository, tokenService),
			UploadAvatar: command.NewUploadAvatarHandler(userRepository, container.Storage),
			UploadCover:  command.NewUploadCoverHandler(userRepository, container.Storage),
			UpdateUser:   command.NewUpdateUserHandler(userRepository),
		},
		Queries: application.Queries{
			GetUser:     query.NewGetUserHandler(userRepository),
			GetAllUsers: query.NewAllUsersHandler(userRepository),
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

func (m *Module) RegisterRoutes() {
	authGroup := m.Router.Group("/api/v1/auth")
	{
		authGroup.POST("/register", m.HttpServer.CreateUser)
		authGroup.POST("/login", m.HttpServer.LoginUser)
		authGroup.GET("/me", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.GetMe)
	}
	userGroup := m.Router.Group("/api/v1/users")
	{
		userGroup.GET("/:id", m.HttpServer.GetUser)
		userGroup.GET("", m.HttpServer.GetAllUsers)
		userGroup.POST("/upload-avatar", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.UploadAvatar)
		userGroup.POST("/upload-cover", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.UploadCover)
		userGroup.PUT("/:id", middlewares.Authenticate(m.Container.JwtService), m.HttpServer.UpdateUser)
	}
}

func (m *Module) GetUserRepository() repository.UserRepository {
	return adapters.NewPgSqlUserRepository(m.DB)
}
