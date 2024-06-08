package command

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/application/dto"
	"github.com/vakhia/artilight/internal/auction/domain/repository"
	"time"
)

type UpdateAuctionCommand struct {
	Request dto.UpdateAuctionRequest
}

type UpdateAuctionHandler struct {
	auctionRepository repository.AuctionRepository
}

func NewUpdateAuctionCommandHandler(auctionRepository repository.AuctionRepository) UpdateAuctionHandler {
	return UpdateAuctionHandler{
		auctionRepository: auctionRepository,
	}
}

func (h *UpdateAuctionHandler) Handle(id uuid.UUID, request dto.UpdateAuctionRequest) error {
	auction, err := h.auctionRepository.FindAuctionById(id)
	if err != nil {
		return err
	}

	if request.StartDate != "" {
		startDate, err := time.Parse(time.RFC3339, request.StartDate)
		if err != nil {
			return err
		}
		auction.StartDate = startDate
	}

	if request.EndDate != "" {
		endDate, err := time.Parse(time.RFC3339, request.EndDate)
		if err != nil {
			return err
		}
		auction.EndDate = endDate
	}

	auction.Status = request.Status
	return h.auctionRepository.Save(auction)
}
