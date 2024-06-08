package application

import (
	"github.com/vakhia/artilight/internal/art/application/command"
	"github.com/vakhia/artilight/internal/art/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateArtCommand        command.CreateArtHandler
	CreateCategoryCommand   command.CreateCategoryHandler
	CreateCollectionCommand command.CreateCollectionHandler
}

type Queries struct {
	AllArts query.AllArtsQuery
	GetArt  query.GetArtHandler
}
