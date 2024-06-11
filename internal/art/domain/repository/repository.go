package repository

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"github.com/vakhia/artilight/internal/common/dtos"
)

type ArtRepository interface {
	GetAllArts(params dtos.PaginationParams, sortParams dtos.SortingParams) ([]aggregate.Art, error)
	GetArtsByCollectionId(collectionId uuid.UUID) ([]aggregate.Art, error)
	FindArtById(id uuid.UUID) (aggregate.Art, error)
	FindArtBySlug(slug string) (aggregate.Art, error)
	Save(art aggregate.Art) error
	GetArtsByOwnerId(id uuid.UUID) ([]aggregate.Art, error)
}

type CategoryRepository interface {
	GetCategories() []entity.Category
	FindCategoryById(id int) (entity.Category, error)
	Save(category entity.Category) error
}

type CollectionRepository interface {
	GetCollections() []entity.Collection
	FindCollectionById(id uuid.UUID) (entity.Collection, error)
	Save(collection entity.Collection) error
}
