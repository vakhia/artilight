package command

import (
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
	"github.com/vakhia/artilight/internal/art/domain/repository"
	"github.com/vakhia/artilight/internal/art/domain/service"
	"github.com/vakhia/artilight/internal/art/domain/valueobject"
)

type CreateArtCommand struct {
	Request dto.CreateArtRequest
}

type CreateArtHandler struct {
	artRepository        repository.ArtRepository
	categoryRepository   repository.CategoryRepository
	userService          service.UserService
	collectionRepository repository.CollectionRepository
}

func NewCreateArtHandler(artRepository repository.ArtRepository, categoryRepository repository.CategoryRepository, userService service.UserService, collectionRepository repository.CollectionRepository) CreateArtHandler {
	return CreateArtHandler{
		artRepository:        artRepository,
		categoryRepository:   categoryRepository,
		userService:          userService,
		collectionRepository: collectionRepository,
	}
}

func (h *CreateArtHandler) Handle(request dto.CreateArtRequest) error {
	owner, err := h.userService.FindOwnerById(request.OwnerId)
	if err != nil {
		return err
	}

	category, err := h.categoryRepository.FindCategoryById(request.CategoryId)
	if err != nil {
		return err
	}

	collection, err := h.collectionRepository.FindCollectionById(request.CollectionId)
	if err != nil {
		return err
	}

	art := aggregate.NewArt(request.Slug, request.Title, request.Description, request.Price, request.MinBid, owner, category, valueobject.Draft, collection)
	return h.artRepository.Save(art)
}
