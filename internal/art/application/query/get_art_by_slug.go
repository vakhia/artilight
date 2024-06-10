package query

import (
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
)

type GetArtBySlugHandler struct {
	readModel GetArtBySlugReadModel
}

func NewGetArtBySlugHandler(readModel GetArtBySlugReadModel) GetArtBySlugHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetArtBySlugHandler{readModel: readModel}
}

type GetArtBySlugReadModel interface {
	FindArtBySlug(slug string) (aggregate.Art, error)
}

func (h GetArtBySlugHandler) Handle(slug string) (dto.ArtResponse, error) {
	art, err := h.readModel.FindArtBySlug(slug)
	if err != nil {
		return dto.ArtResponse{}, err
	}

	return mapArtToArtResponse(art), err
}
