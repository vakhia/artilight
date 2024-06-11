package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateBidRequest struct {
	Amount    float64   `json:"amount"`
	ItemID    uuid.UUID `json:"item_id"`
	AuctionID uuid.UUID `json:"auction_id"`
	BidderID  uuid.UUID `json:"bidder_id"`
}

type BidResponse struct {
	Id     uuid.UUID      `json:"id"`
	Amount float64        `json:"amount"`
	Time   time.Time      `json:"time"`
	Item   ItemResponse   `json:"item"`
	Bidder BidderResponse `json:"bidder"`
}

type BidderResponse struct {
	Id       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Avatar   string    `json:"avatar"`
}
