package container

import (
	artModule "github.com/vakhia/artilight/internal/art/domain/repository"
	"github.com/vakhia/artilight/internal/common/fileuploader"
	"github.com/vakhia/artilight/internal/user/domain/repository"
	"github.com/vakhia/artilight/pkg/token"
)

type Container struct {
	JwtService     token.IJwtService
	UserRepository repository.UserRepository
	Storage        fileuploader.IStorage
	ArtRepository  artModule.ArtRepository
}

func NewContainer() *Container {
	return &Container{}
}
