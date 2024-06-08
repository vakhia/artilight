package application

import (
	"github.com/vakhia/artilight/internal/user/application/command"
	"github.com/vakhia/artilight/internal/user/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser   command.CreateUserHandler
	LoginUser    command.LoginUserHandler
	UploadAvatar command.UploadAvatarHandler
	UploadCover  command.UploadCoverHandler
	UpdateUser   command.UpdateUserHandler
}

type Queries struct {
	GetUser query.GetUserHandler
}
