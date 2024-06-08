package query

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
)

type GetArtHandler struct {
	readModel GetUserReadModel
}

func NewGetArtHandler(readModel GetUserReadModel) GetArtHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetArtHandler{readModel: readModel}
}

type GetUserReadModel interface {
	FindArtById(id uuid.UUID) (aggregate.Art, error)
}

func (h GetArtHandler) Handle(id uuid.UUID) (dto.ArtResponse, error) {
	art, err := h.readModel.FindArtById(id)
	if err != nil {
		return dto.ArtResponse{}, err
	}

	return mapArtToArtResponse(art), err
}
