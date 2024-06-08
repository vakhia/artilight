package aggregate

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"github.com/vakhia/artilight/internal/art/domain/valueobject"
)

type Art struct {
	Id           uuid.UUID
	Slug         string
	Title        string
	Description  string
	Price        float64
	Status       valueobject.ArtStatus
	OwnerId      uuid.UUID         `json:"-"`
	Owner        entity.Owner      `json:"owner,omitempty"`
	CollectionId uuid.UUID         `json:"-"`
	Collection   entity.Collection `json:"collection,omitempty"`
	CategoryId   int               `json:"-"`
	Category     entity.Category   `json:"category,omitempty"`
}

func NewArt(slug, title, description string, price float64, owner entity.Owner, category entity.Category, status valueobject.ArtStatus, collection entity.Collection) Art {
	return Art{
		Id:           uuid.New(),
		Slug:         slug,
		Title:        title,
		Description:  description,
		Price:        price,
		Status:       status,
		Owner:        owner,
		OwnerId:      owner.Id,
		CollectionId: collection.Id,
		Collection:   collection,
		Category:     category,
		CategoryId:   category.Id,
	}
}
