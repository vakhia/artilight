package query

import (
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/art/domain/entity"
)

type AllCategoriesQuery struct {
	readModel AllCategoriesReadModel
}

func NewAllCategoriesQuery(readModel AllCategoriesReadModel) AllCategoriesQuery {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllCategoriesQuery{readModel: readModel}
}

type AllCategoriesReadModel interface {
	GetCategories() []entity.Category
}

func (h AllCategoriesQuery) Handle() ([]dto.CategoryResponse, error) {
	categories := h.readModel.GetCategories()

	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponse := mapCategoryToCategoryResponse(category)
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses, nil
}

func mapCategoryToCategoryResponse(category entity.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id:          category.Id,
		Title:       category.Title,
		Description: category.Description,
	}
}
