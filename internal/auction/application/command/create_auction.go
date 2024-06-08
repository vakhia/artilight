package command

import (
	"github.com/vakhia/artilight/internal/auction/application/dto"
	"github.com/vakhia/artilight/internal/auction/domain/aggregate"
	"github.com/vakhia/artilight/internal/auction/domain/repository"
	"github.com/vakhia/artilight/internal/auction/domain/service"
	"github.com/vakhia/artilight/internal/auction/domain/valueobject"
	"time"
)

type CreateAuctionCommand struct {
	Request dto.CreateAuctionRequest
}

type CreateAuctionHandler struct {
	auctionRepository repository.AuctionRepository
	userService       service.UserService
	itemService       service.ItemService
}

func NewCreateAuctionCommandHandler(auctionRepository repository.AuctionRepository, userService service.UserService, itemService service.ItemService) CreateAuctionHandler {
	return CreateAuctionHandler{
		auctionRepository: auctionRepository,
		userService:       userService,
		itemService:       itemService,
	}
}

func (h *CreateAuctionHandler) Handle(request dto.CreateAuctionRequest) error {
	item, err := h.itemService.FindItemById(request.ItemId)
	if err != nil {
		return err
	}

	startDate, err := time.Parse(time.RFC3339, request.StartDate)
	if err != nil {
		return err
	}

	endDate, err := time.Parse(time.RFC3339, request.EndDate)
	if err != nil {
		return err
	}

	auction := aggregate.NewAuction(valueobject.Created, valueobject.Bid, startDate, endDate, item)
	return h.auctionRepository.Save(auction)
}
