package query

import (
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/entity"
)

type AllCollectionsQuery struct {
	readModel AllCollectionReadModel
}

func NewAllCollectionQuery(readModel AllCollectionReadModel) AllCollectionsQuery {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllCollectionsQuery{readModel: readModel}
}

type AllCollectionReadModel interface {
	GetCollections() []entity.Collection
}

func (h AllCollectionsQuery) Handle() ([]dto.CollectionResponse, error) {
	collections := h.readModel.GetCollections()

	var collectionResponses []dto.CollectionResponse
	for _, collection := range collections {
		collectionResponse := mapCollectionToCollectionResponse(collection)
		collectionResponses = append(collectionResponses, collectionResponse)
	}

	return collectionResponses, nil
}

func mapCollectionToCollectionResponse(collection entity.Collection) dto.CollectionResponse {
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
	}
}
