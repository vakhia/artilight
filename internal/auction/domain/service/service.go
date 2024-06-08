package service

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/entity"
)

type UserService interface {
	FindBidderById(userID uuid.UUID) (entity.Bidder, error)
}

type ItemService interface {
	FindItemById(artId uuid.UUID) (entity.Item, error)
}
