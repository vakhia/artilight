package repositories

import (
	"github.com/vakhia/artilight/internal/domain"
	"gorm.io/gorm"
)

type IArtRepository interface {
	GetAllArts() ([]domain.Art, error)
	GetArtById(id int) (domain.Art, error)
	CreateArt(art domain.Art) (domain.Art, error)
	UpdateArt(art domain.Art) (domain.Art, error)
	DeleteArt(id int) error
}

type ArtRepository struct {
	db *gorm.DB
}

func NewArtRepository(db *gorm.DB) *ArtRepository {
	return &ArtRepository{db: db}
}

func (r *ArtRepository) GetAllArts() ([]domain.Art, error) {
	var arts []domain.Art
	err := r.db.Preload("Owner").Preload("Category").Find(&arts).Error
	if err != nil {
		return nil, err
	}
	return arts, nil
}

func (r *ArtRepository) GetArtById(id int) (domain.Art, error) {
	var art domain.Art
	err := r.db.Preload("Owner").Preload("Category").First(&art, id).Error
	if err != nil {
		return domain.Art{}, err
	}
	return art, nil
}

func (r *ArtRepository) CreateArt(art domain.Art) (domain.Art, error) {
	err := r.db.Create(&art).Error
	if err != nil {
		return domain.Art{}, err
	}
	return art, nil
}

func (r *ArtRepository) UpdateArt(art domain.Art) (domain.Art, error) {
	err := r.db.Save(&art).Error
	if err != nil {
		return domain.Art{}, err
	}
	return art, nil
}

func (r *ArtRepository) DeleteArt(id int) error {
	err := r.db.Delete(&domain.Art{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
