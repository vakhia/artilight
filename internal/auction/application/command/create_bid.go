package command

import (
	"github.com/vakhia/artilight/internal/auction/application/dto"
	"github.com/vakhia/artilight/internal/auction/domain/aggregate"
	"github.com/vakhia/artilight/internal/auction/domain/repository"
	"github.com/vakhia/artilight/internal/auction/domain/service"
)

type CreateBidHandler struct {
	bidRepository     repository.BidRepository
	userService       service.UserService
	itemService       service.ItemService
	auctionRepository repository.AuctionRepository
}

func NewCreateBidCommandHandler(bidRepository repository.BidRepository, userService service.UserService, itemService service.ItemService, auctionRepository repository.AuctionRepository) CreateBidHandler {
	return CreateBidHandler{
		bidRepository:     bidRepository,
		userService:       userService,
		itemService:       itemService,
		auctionRepository: auctionRepository,
	}
}

func (h *CreateBidHandler) Handle(request dto.CreateBidRequest) error {
	item, err := h.itemService.FindItemById(request.ItemID)
	if err != nil {
		return err
	}

	bidder, err := h.userService.FindBidderById(request.BidderID)
	if err != nil {
		return err
	}

	auction, err := h.auctionRepository.FindAuctionById(request.AuctionID)
	if err != nil {
		return err
	}

	bid := aggregate.NewBid(request.Amount, item, bidder, auction)
	return h.bidRepository.Save(bid)
}
