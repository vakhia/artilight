package dto

import (
	"github.com/google/uuid"
)

type CreateCollectionRequest struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	OwnerId     uuid.UUID `json:"owner_id"`
}

type CollectionResponse struct {
	Id          uuid.UUID     `json:"id"`
	Slug        string        `json:"slug"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Owner       OwnerResponse `json:"owner"`
}
