package query

//type GetArtHandler struct {
//	readModel GetArtReadModel
//}
//
//func NewGetArt(readModel GetArtReadModel) GetArtHandler {
//	if readModel == nil {
//		panic("nil readModel")
//	}
//
//	return GetArtHandler{readModel: readModel}
//}
//
//type GetArtReadModel interface {
//	GetArt(ctx context.Context) (art , err error)
//}
//
//func (h GetArtHandler) Handle(ctx context.Context) (art domain.Art, err error) {
//	return h.readModel.GetArt(ctx)
//}
