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
	MinBid       float64
	Status       valueobject.ArtStatus
	OwnerId      uuid.UUID         `json:"-"`
	Owner        entity.Owner      `json:"owner,omitempty"`
	CollectionId uuid.UUID         `json:"-"`
	Collection   entity.Collection `json:"collection,omitempty"`
	CategoryId   int               `json:"-"`
	Category     entity.Category   `json:"category,omitempty"`
	Auctions     []entity.Auction  `gorm:"foreignKey:ItemID"` // Use 'ItemID' as the foreign key
	Images       []entity.ArtImage `gorm:"foreignKey:ArtId"`  // Use 'ArtId' as the foreign key
	Tags         []entity.Tag      `gorm:"many2many:arts_tags;"`
}

func NewArt(slug, title, description string, price float64, minBid float64, owner entity.Owner, category entity.Category, status valueobject.ArtStatus, collection entity.Collection) Art {
	return Art{
		Id:           uuid.New(),
		Slug:         slug,
		Title:        title,
		Description:  description,
		Price:        price,
		MinBid:       minBid,
		Status:       status,
		Owner:        owner,
		OwnerId:      owner.Id,
		CollectionId: collection.Id,
		Collection:   collection,
		Category:     category,
		CategoryId:   category.Id,
	}
}

func (a *Art) AddImage(image entity.ArtImage) {
	for _, img := range a.Images {
		if img.Url == image.Url {
			return
		}
	}

	a.Images = append(a.Images, image)
}

func (a *Art) AddTag(tag entity.Tag) {
	for _, t := range a.Tags {
		if t.Id == tag.Id {
			return
		}
	}

	a.Tags = append(a.Tags, tag)
}
