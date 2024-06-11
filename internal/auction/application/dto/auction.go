package dto

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/valueobject"
	"time"
)

type CreateAuctionRequest struct {
	ItemId    uuid.UUID `json:"item_id"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
}

type UpdateAuctionRequest struct {
	Status    valueobject.AuctionStatus `json:"status"`
	StartDate string                    `json:"start_date"`
	EndDate   string                    `json:"end_date"`
}

type AuctionResponse struct {
	Id           uuid.UUID     `json:"id"`
	Status       string        `json:"status"`
	StartDate    time.Time     `json:"start_date"`
	EndDate      time.Time     `json:"end_date"`
	InitialPrice float64       `json:"initial_price"`
	CurrentPrice float64       `json:"current_price"`
	Item         ItemResponse  `json:"item"`
	Bids         []BidResponse `json:"bids"`
}

type ItemResponse struct {
	Id    uuid.UUID `json:"id"`
	Slug  string    `json:"slug"`
	Title string    `json:"title"`
}
