package query

import (
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
	"github.com/vakhia/artilight/internal/common/dtos"
)

type AllArtsQuery struct {
	readModel AllArtsReadModel
}

func NewAllArtsQuery(readModel AllArtsReadModel) AllArtsQuery {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllArtsQuery{readModel: readModel}
}

type AllArtsReadModel interface {
	GetAllArts(params dtos.PaginationParams, sortingParams dtos.SortingParams) ([]aggregate.Art, error)
}

func (h AllArtsQuery) Handle(params dtos.PaginationParams, sortingParams dtos.SortingParams) ([]dto.ArtResponse, error) {
	arts, err := h.readModel.GetAllArts(params, sortingParams)
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

func mapArtToArtResponse(art aggregate.Art) dto.ArtResponse {
	return dto.ArtResponse{
		ID:          art.Id,
		Slug:        art.Slug,
		Title:       art.Title,
		Description: art.Description,
		Price:       art.Price,
		Status:      art.Status.String(),
		Owner: dto.OwnerResponse{
			Id:       art.Owner.Id,
			FullName: art.Owner.GetFullName(),
			Avatar:   art.Owner.Avatar,
		},
		Category: dto.CategoryResponse{
			Id:          art.Category.Id,
			Title:       art.Category.Title,
			Description: art.Category.Description},
		Collection: dto.CollectionResponse{
			Id:          art.Collection.Id,
			Title:       art.Collection.Title,
			Slug:        art.Collection.Slug,
			Description: art.Collection.Description,
			Owner: dto.OwnerResponse{
				Id:       art.Collection.Owner.Id,
				FullName: art.Collection.Owner.GetFullName(),
				Avatar:   art.Collection.Owner.Avatar,
			},
		},
	}
}
