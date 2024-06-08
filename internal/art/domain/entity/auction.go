package entity

import "github.com/google/uuid"

type Auction struct {
	Id        uuid.UUID
	ItemID    uuid.UUID `gorm:"index;foreignKey:ItemID;references:Id"`
	Status    string
	StartDate string
	EndDate   string
	Bids      []Bid `gorm:"foreignKey:AuctionID"`
}

func (a Auction) TableName() string {
	return "auctions"
}

type Bid struct {
	Id        uuid.UUID
	Amount    float64
	BidderID  uuid.UUID `gorm:"index;foreignKey:BidderID;references:ID"`
	Bidder    Owner     `gorm:"foreignKey:BidderID"`
	AuctionID uuid.UUID `gorm:"index;foreignKey:AuctionID;references:ID"`
}

func (b Bid) TableName() string {
	return "bids"
}
