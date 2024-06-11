package query

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
)

type AllUsersArtsQuery struct {
	readModel AllUsersArtsReadModel
}

func NewAllUsersArtsQuery(readModel AllUsersArtsReadModel) AllUsersArtsQuery {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllUsersArtsQuery{readModel: readModel}
}

type AllUsersArtsReadModel interface {
	GetArtsByOwnerId(uuid uuid.UUID) ([]aggregate.Art, error)
}

func (h AllUsersArtsQuery) Handle(userId uuid.UUID) ([]dto.ArtResponse, error) {
	arts, err := h.readModel.GetArtsByOwnerId(userId)
	if err != nil {
		return nil, err
	}

	var artResponses []dto.ArtResponse
	for _, art := range arts {
		artResponse := mapArtToArtResponse(art)
		artResponses = append(artResponses, artResponse)
	}

	return artResponses, nil
}
