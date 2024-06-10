package query

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
	"github.com/vakhia/artilight/internal/art/domain/entity"
)

type AllCollectionsQuery struct {
	readModel    AllCollectionReadModel
	artReadModel AllArtsByCollectionIdReadModel
}

func NewAllCollectionQuery(readModel AllCollectionReadModel, artReadModel AllArtsByCollectionIdReadModel) AllCollectionsQuery {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllCollectionsQuery{readModel: readModel, artReadModel: artReadModel}
}

type AllCollectionReadModel interface {
	GetCollections() []entity.Collection
}

type AllArtsByCollectionIdReadModel interface {
	GetArtsByCollectionId(collectionId uuid.UUID) ([]aggregate.Art, error)
}

func (h AllCollectionsQuery) Handle() ([]dto.CollectionResponse, error) {
	collections := h.readModel.GetCollections()

	var collectionResponses []dto.CollectionResponse
	for _, collection := range collections {
		arts, _ := h.artReadModel.GetArtsByCollectionId(collection.Id)
		collectionResponse := mapCollectionToCollectionResponse(collection, arts)
		collectionResponses = append(collectionResponses, collectionResponse)
	}

	return collectionResponses, nil
}

func mapCollectionToCollectionResponse(collection entity.Collection, arts []aggregate.Art) dto.CollectionResponse {
	var artResponses []dto.ArtResponse
	for _, art := range arts {
		artResponses = append(artResponses, mapArtToArtResponse(art))
	}

	return dto.CollectionResponse{
		Id:          collection.Id,
		Title:       collection.Title,
		Description: collection.Description,
		Slug:        collection.Slug,
		Owner: dto.OwnerResponse{
			Id:       collection.Owner.Id,
			FullName: collection.Owner.GetFullName(),
			Avatar:   collection.Owner.Avatar,
		},
		Arts: artResponses,
	}
}
