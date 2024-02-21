package services

import (
	"github.com/vakhia/artilight/internal/common/dto"
	"github.com/vakhia/artilight/internal/domain"
	"github.com/vakhia/artilight/internal/repositories"
)

type ICategoryService interface {
	GetAllCategories() ([]dto.CategoryResponse, error)
	GetCategoryById(id int) (dto.CategoryResponse, error)
	CreateCategory(category dto.CategoryCreateRequest) (dto.CategoryResponse, error)
	UpdateCategory(category dto.CategoryUpdateRequest) (dto.CategoryResponse, error)
	DeleteCategory(id int) error
}

type CategoryService struct {
	categoryRepo repositories.ICategoryRepository
}

func NewCategoryService(categoryRepo repositories.ICategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (c *CategoryService) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := c.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.CategoryResponse{
			Id:          category.Id,
			Title:       category.Title,
			Description: category.Description,
		})
	}
	return categoryResponses, nil
}

func (c *CategoryService) GetCategoryById(id int) (dto.CategoryResponse, error) {
	category, err := c.categoryRepo.GetCategoryById(id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	return dto.CategoryResponse{
		Id:          category.Id,
		Title:       category.Title,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategory(category dto.CategoryCreateRequest) (dto.CategoryResponse, error) {
	categoryDomain := domain.Category{
		Title:       category.Title,
		Description: category.Description,
	}
	createdCategory, err := c.categoryRepo.CreateCategory(categoryDomain)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	return dto.CategoryResponse{
		Id:          createdCategory.Id,
		Title:       createdCategory.Title,
		Description: createdCategory.Description,
	}, nil
}

func (c *CategoryService) UpdateCategory(category dto.CategoryUpdateRequest) (dto.CategoryResponse, error) {
	_, err := c.categoryRepo.GetCategoryById(category.Id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	categoryDomain := domain.Category{
		Title:       category.Title,
		Description: category.Description,
	}
	categoryDomain.Id = category.Id
	updatedCategory, err := c.categoryRepo.UpdateCategory(categoryDomain)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	return dto.CategoryResponse{
		Id:          updatedCategory.Id,
		Title:       updatedCategory.Title,
		Description: updatedCategory.Description,
	}, nil
}

func (c *CategoryService) DeleteCategory(id int) error {
	return c.categoryRepo.DeleteCategory(id)
}
