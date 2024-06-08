package query

//type AllArtsQuery struct {
//	readModel AllArtsReadModel
//}
//
//func NewAllArtsQuery(readModel AllArtsReadModel) AllArtsQuery {
//	if readModel == nil {
//		panic("nil readModel")
//	}
//
//	return AllArtsQuery{readModel: readModel}
//}
//
//type AllArtsReadModel interface {
//	GetAllArts(pageSize, pageNumber int, sortBy, sortOrder string) ([]aggregate.Art, error)
//}
//
//func (h AllArtsQuery) Handle(pageSize, pageNumber int, sortBy, sortOrder string) ([]dto.ArtResponse, error) {
//	arts, err := h.readModel.GetAllArts(pageSize, pageNumber, sortBy, sortOrder)
//	if err != nil {
//		return nil, err
//	}
//
//	var artResponses []dto.ArtResponse
//	for _, art := range arts {
//		artResponse := mapArtToArtResponse(art)
//		artResponses = append(artResponses, artResponse)
//	}
//
//	return artResponses, nil
//}
//
//func mapArtToArtResponse(art aggregate.Art) dto.ArtResponse {
//	return dto.ArtResponse{
//		ID:          art.Id,
//		Slug:        art.Slug,
//		Title:       art.Title,
//		Description: art.Description,
//		Price:       art.Price,
//		Status:      art.Status.String(),
//		Owner:       art.Owner,
//		Category:    art.Category,
//	}
//}
