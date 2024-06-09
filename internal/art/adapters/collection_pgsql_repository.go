package adapters

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgSqlCollectionRepository struct {
	db *gorm.DB
}

func NewPgSqlCollectionRepository(db *gorm.DB) *PgSqlCollectionRepository {
	if db == nil {
		panic("missing db")
	}

	return &PgSqlCollectionRepository{db: db}
}

func (r *PgSqlCollectionRepository) GetCollections() []entity.Collection {
	var collections []entity.Collection
	r.db.Preload("Owner").Find(&collections)
	return collections
}

func (r *PgSqlCollectionRepository) FindCollectionById(id uuid.UUID) (entity.Collection, error) {
	var collection entity.Collection
	result := r.db.Where("id = ?", id).First(&collection)
	return collection, result.Error
}

func (r *PgSqlCollectionRepository) Save(collection entity.Collection) error {
	result := r.db.Omit(clause.Associations).Create(&collection)
	return result.Error
}
