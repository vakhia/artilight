package dto

import (
	"github.com/google/uuid"
)

type CreateArtRequest struct {
	Slug         string    `json:"slug" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Price        float64   `json:"price" binding:"required"`
	Status       string    `json:"status"`
	CategoryId   int       `json:"category_id" binding:"required"`
	OwnerId      uuid.UUID `json:"owner_id" binding:"required"`
	CollectionId uuid.UUID `json:"collection_id" binding:"required"`
}

type ArtResponse struct {
	ID          uuid.UUID          `json:"id"`
	Slug        string             `json:"slug"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Status      string             `json:"status"`
	Owner       OwnerResponse      `json:"owner,omitempty"`
	Category    CategoryResponse   `json:"category"`
	Collection  CollectionResponse `json:"collection"`
}
