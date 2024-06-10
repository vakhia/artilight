package adapters

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/aggregate"
	"github.com/vakhia/artilight/internal/common/dtos"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgSqlArtRepository struct {
	db *gorm.DB
}

func NewPgSqlArtRepository(db *gorm.DB) *PgSqlArtRepository {
	if db == nil {
		panic("missing db")
	}

	return &PgSqlArtRepository{db: db}
}

func (r *PgSqlArtRepository) Save(art aggregate.Art) error {
	if err := r.db.Omit(clause.Associations).Save(&art).Error; err != nil {
		return err
	}

	if err := r.db.Model(&art).Association("Images").Append(art.Images); err != nil {
		return err
	}

	return nil
}

func (r *PgSqlArtRepository) FindArtById(id uuid.UUID) (aggregate.Art, error) {
	var art aggregate.Art
	result := r.db.Preload("Category").Preload("Collection.Owner").Preload("Images").Preload("Auctions.Bids.Bidder").Preload("Owner").Where("id = ?", id).First(&art)
	return art, result.Error
}

func (r *PgSqlArtRepository) FindArtBySlug(slug string) (aggregate.Art, error) {
	var art aggregate.Art
	result := r.db.Preload("Category").Preload("Collection.Owner").Preload("Images").Preload("Auctions.Bids.Bidder").Preload("Owner").Where("slug = ?", slug).First(&art)
	return art, result.Error
}

func (r *PgSqlArtRepository) GetAllArts(params dtos.PaginationParams, sortParams dtos.SortingParams) ([]aggregate.Art, error) {
	var arts []aggregate.Art
	query := r.db.Preload("Category").Preload("Collection.Owner").Preload("Images").Preload("Auctions.Bids.Bidder").Preload("Owner").Preload(clause.Associations)

	// Sorting
	if sortParams.SortBy != "" {
		if sortParams.SortOrder == "desc" {
			query = query.Order(fmt.Sprintf("%s desc", sortParams.SortBy))
		} else {
			query = query.Order(fmt.Sprintf("%s asc", sortParams.SortBy))
		}
	}

	// Pagination
	if params.PageNumber > 0 && params.PageSize > 0 {
		offset := (params.PageNumber - 1) * params.PageSize
		query = query.Offset(offset).Limit(params.PageSize)
	}

	// Execute the query
	if err := query.Find(&arts).Error; err != nil {
		return nil, err
	}
	return arts, nil
}
