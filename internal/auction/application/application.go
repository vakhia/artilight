package application

import (
	"github.com/vakhia/artilight/internal/auction/application/command"
	"github.com/vakhia/artilight/internal/auction/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateAuctionCommand command.CreateAuctionHandler
	UpdateAuctionCommand command.UpdateAuctionHandler
	CreateBidCommand     command.CreateBidHandler
}

type Queries struct {
	GetAuctionQuery query.GetAuctionHandler
}
