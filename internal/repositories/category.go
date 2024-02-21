package repositories

import (
	"github.com/vakhia/artilight/internal/domain"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	GetAllCategories() ([]domain.Category, error)
	GetCategoryById(id int) (domain.Category, error)
	CreateCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(category domain.Category) (domain.Category, error)
	DeleteCategory(id int) error
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryById(id int) (domain.Category, error) {
	var category domain.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) CreateCategory(category domain.Category) (domain.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(category domain.Category) (domain.Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	err := r.db.Delete(&domain.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
