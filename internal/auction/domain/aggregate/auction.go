package aggregate

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/entity"
	"github.com/vakhia/artilight/internal/auction/domain/valueobject"
	"time"
)

type Auction struct {
	Id        uuid.UUID
	Status    valueobject.AuctionStatus
	Type      valueobject.AuctionType
	StartDate time.Time
	EndDate   time.Time
	ItemId    uuid.UUID   `json:"-"`
	Item      entity.Item `json:"item,omitempty"`
	Bids      []Bid       `json:"bids,omitempty"`
}

func NewAuction(status valueobject.AuctionStatus, auctionType valueobject.AuctionType, startDate, endDate time.Time, item entity.Item) Auction {
	return Auction{
		Id:        uuid.New(),
		Status:    status,
		Type:      auctionType,
		StartDate: startDate,
		EndDate:   endDate,
		ItemId:    item.Id,
		Item:      item,
	}
}
