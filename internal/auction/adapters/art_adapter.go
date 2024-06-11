package adapters

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/repository"
	"github.com/vakhia/artilight/internal/auction/domain/entity"
)

type ArtAdapter struct {
	artRepository repository.ArtRepository
}

func NewArtAdapter(artRepository repository.ArtRepository) *ArtAdapter {
	return &ArtAdapter{artRepository: artRepository}
}

func (a *ArtAdapter) FindItemById(artId uuid.UUID) (entity.Item, error) {
	art, err := a.artRepository.FindArtById(artId)
	if err != nil {
		return entity.Item{}, err
	}

	return entity.Item{
		Id:    art.Id,
		Slug:  art.Slug,
		Title: art.Title,
		Price: art.Price,
	}, nil
}
