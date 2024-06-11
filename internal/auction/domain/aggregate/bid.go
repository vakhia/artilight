package aggregate

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/entity"
	"time"
)

type Bid struct {
	Id        uuid.UUID
	Amount    float64
	Time      time.Time
	ItemID    uuid.UUID
	Item      entity.Item
	BidderID  uuid.UUID
	Bidder    entity.Bidder
	AuctionID uuid.UUID
	Auction   Auction
}

func NewBid(amount float64, item entity.Item, bidder entity.Bidder, auction Auction) Bid {
	return Bid{
		Id:        uuid.New(),
		Amount:    amount,
		Time:      time.Now(),
		Item:      item,
		ItemID:    item.Id,
		Bidder:    bidder,
		BidderID:  bidder.Id,
		Auction:   auction,
		AuctionID: auction.Id,
	}
}
