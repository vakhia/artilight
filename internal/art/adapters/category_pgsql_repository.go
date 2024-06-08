package adapters

import (
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"gorm.io/gorm"
)

type PgSqlCategoryRepository struct {
	db *gorm.DB
}

func NewPgSqlCategoryRepository(db *gorm.DB) *PgSqlCategoryRepository {
	if db == nil {
		panic("missing db")
	}

	return &PgSqlCategoryRepository{db: db}
}

func (r *PgSqlCategoryRepository) GetCategories() []entity.Category {
	var categories []entity.Category
	r.db.Find(&categories)
	return categories
}

func (r *PgSqlCategoryRepository) FindCategoryById(id int) (entity.Category, error) {
	var category entity.Category
	result := r.db.Where("id = ?", id).First(&category)
	return category, result.Error
}

func (r *PgSqlCategoryRepository) Save(category entity.Category) error {
	result := r.db.Save(&category)
	return result.Error
}
