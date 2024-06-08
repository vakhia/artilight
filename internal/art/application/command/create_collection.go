package command

import (
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"github.com/vakhia/artilight/internal/art/domain/repository"
	"github.com/vakhia/artilight/internal/art/domain/service"
)

type CreateCollectionCommand struct {
	Request dto.CreateArtRequest
}

type CreateCollectionHandler struct {
	collectionRepository repository.CollectionRepository
	userService          service.UserService
}

func NewCollectionHandler(collectionRepository repository.CollectionRepository, userService service.UserService) CreateCollectionHandler {
	return CreateCollectionHandler{
		collectionRepository: collectionRepository,
		userService:          userService,
	}
}

func (h *CreateCollectionHandler) Handle(request dto.CreateCollectionRequest) error {
	owner, err := h.userService.FindOwnerById(request.OwnerId)
	if err != nil {
		return err
	}

	collection := entity.NewCollection(request.Slug, request.Title, request.Description, owner)
	return h.collectionRepository.Save(collection)
}
