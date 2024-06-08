package command

import (
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"github.com/vakhia/artilight/internal/art/domain/repository"
)

type CreateCategoryCommand struct {
	Request dto.CreateCategoryRequest
}

type CreateCategoryHandler struct {
	categoryRepository repository.CategoryRepository
}

func NewCreateCategoryHandler(categoryRepository repository.CategoryRepository) CreateCategoryHandler {
	return CreateCategoryHandler{
		categoryRepository: categoryRepository,
	}
}

func (h *CreateCategoryHandler) Handle(request dto.CreateCategoryRequest) error {
	category := entity.NewCategory(request.Title, request.Description)
	return h.categoryRepository.Save(category)
}
