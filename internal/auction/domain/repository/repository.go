package repository

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/aggregate"
)

type AuctionRepository interface {
	GetAllAuctions() ([]aggregate.Auction, error)
	FindAuctionById(id uuid.UUID) (aggregate.Auction, error)
	Save(auction aggregate.Auction) error
}

type BidRepository interface {
	GetAllBids() ([]aggregate.Bid, error)
	GetBidsByAuctionId(auctionId uuid.UUID) ([]aggregate.Bid, error)
	Save(bid aggregate.Bid) error
}
