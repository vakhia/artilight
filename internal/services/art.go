package services

import (
	"github.com/vakhia/artilight/internal/common/dto"
	"github.com/vakhia/artilight/internal/domain"
	"github.com/vakhia/artilight/internal/repositories"
)

type IArtService interface {
	GetAllArts() ([]dto.ArtResponse, error)
	GetArtById(id int) (dto.ArtResponse, error)
	CreateArt(art dto.ArtCreateRequest) (dto.ArtResponse, error)
	UpdateArt(art dto.ArtUpdateRequest) (dto.ArtResponse, error)
	DeleteArt(id int) error
}

type ArtService struct {
	artRepo repositories.IArtRepository
}

func NewArtService(artRepo repositories.IArtRepository) *ArtService {
	return &ArtService{artRepo: artRepo}
}

func (a *ArtService) GetAllArts() ([]dto.ArtResponse, error) {
	arts, err := a.artRepo.GetAllArts()
	if err != nil {
		return nil, err
	}
	var artResponses []dto.ArtResponse
	for _, art := range arts {
		artResponses = append(artResponses, dto.ArtResponse{
			Id:          art.Id,
			Title:       art.Title,
			Description: art.Description,
			Price:       art.Price,
			Category: &dto.CategoryResponse{
				Id:          art.Category.Id,
				Title:       art.Category.Title,
				Description: art.Category.Description,
			},
			Owner: &dto.UserResponse{
				Id:    art.Owner.Id,
				Name:  art.Owner.Name,
				Email: art.Owner.Email,
			},
		})
	}
	return artResponses, nil
}

func (a *ArtService) GetArtById(id int) (dto.ArtResponse, error) {
	art, err := a.artRepo.GetArtById(id)
	if err != nil {
		return dto.ArtResponse{}, err
	}
	return dto.ArtResponse{
		Id:          art.Id,
		Title:       art.Title,
		Description: art.Description,
		Price:       art.Price,
		Category: &dto.CategoryResponse{
			Id:          art.Category.Id,
			Title:       art.Category.Title,
			Description: art.Category.Description,
		},
		Owner: &dto.UserResponse{
			Id:    art.Owner.Id,
			Name:  art.Owner.Name,
			Email: art.Owner.Email,
		},
	}, nil
}

func (a *ArtService) CreateArt(art dto.ArtCreateRequest) (dto.ArtResponse, error) {
	artDomain := domain.Art{
		Title:       art.Title,
		Description: art.Description,
		Price:       art.Price,
		CategoryId:  art.CategoryId,
		OwnerId:     art.OwnerId,
	}
	createdArt, err := a.artRepo.CreateArt(artDomain)
	if err != nil {
		return dto.ArtResponse{}, err
	}
	return dto.ArtResponse{
		Id:          createdArt.Id,
		Title:       createdArt.Title,
		Description: createdArt.Description,
		Price:       createdArt.Price,
	}, nil
}

func (a *ArtService) UpdateArt(art dto.ArtUpdateRequest) (dto.ArtResponse, error) {
	artDomain := domain.Art{
		Title:       art.Title,
		Description: art.Description,
		Price:       art.Price,
		CategoryId:  art.CategoryId,
		OwnerId:     art.OwnerId,
	}
	artDomain.Id = art.Id
	updatedArt, err := a.artRepo.UpdateArt(artDomain)
	if err != nil {
		return dto.ArtResponse{}, err
	}
	return dto.ArtResponse{
		Id:          updatedArt.Id,
		Title:       updatedArt.Title,
		Description: updatedArt.Description,
		Price:       updatedArt.Price,
	}, nil
}

func (a *ArtService) DeleteArt(id int) error {
	err := a.artRepo.DeleteArt(id)
	if err != nil {
		return err
	}
	return nil
}
