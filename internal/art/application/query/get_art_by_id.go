package query

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
)

type GetArtByIdHandler struct {
	readModel GetArtByIdReadModel
}

func NewGetArtByIdHandler(readModel GetArtByIdReadModel) GetArtByIdHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetArtByIdHandler{readModel: readModel}
}

type GetArtByIdReadModel interface {
	FindArtById(id uuid.UUID) (aggregate.Art, error)
}

func (h GetArtByIdHandler) Handle(id uuid.UUID) (dto.ArtResponse, error) {
	art, err := h.readModel.FindArtById(id)
	if err != nil {
		return dto.ArtResponse{}, err
	}

	return mapArtToArtResponse(art), err
}
