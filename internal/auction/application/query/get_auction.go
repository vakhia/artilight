package query

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/application/dto"
	"github.com/vakhia/artilight/internal/auction/domain/aggregate"
)

type GetAuctionHandler struct {
	readModel GetAuctionReadModel
}

func NewGetAuctionHandler(readModel GetAuctionReadModel) GetAuctionHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetAuctionHandler{readModel: readModel}
}

type GetAuctionReadModel interface {
	FindAuctionById(id uuid.UUID) (auction aggregate.Auction, err error)
}

func (h GetAuctionHandler) Handle(id uuid.UUID) (dto.AuctionResponse, error) {
	auction, err := h.readModel.FindAuctionById(id)
	if err != nil {
		return dto.AuctionResponse{}, err
	}

	return mapAuctionToResponse(auction), err
}

func mapAuctionToResponse(auction aggregate.Auction) dto.AuctionResponse {
	return dto.AuctionResponse{
		Id:           auction.Id,
		Status:       auction.Status.String(),
		StartDate:    auction.StartDate,
		EndDate:      auction.EndDate,
		InitialPrice: auction.InitialPrice,
		CurrentPrice: auction.CurrentPrice,
		Item: dto.ItemResponse{
			Id:    auction.Item.Id,
			Slug:  auction.Item.Slug,
			Title: auction.Item.Title,
		},
		Bids: mapBidsToResponse(auction.Bids),
	}
}

func mapBidsToResponse(bids []aggregate.Bid) []dto.BidResponse {
	var response []dto.BidResponse
	for _, bid := range bids {
		response = append(response, dto.BidResponse{
			Id:     bid.Id,
			Amount: bid.Amount,
			Time:   bid.Time,
			Item: dto.ItemResponse{
				Id:    bid.Item.Id,
				Slug:  bid.Item.Slug,
				Title: bid.Item.Title,
			},
			Bidder: dto.BidderResponse{
				Id:       bid.Bidder.Id,
				FullName: bid.Bidder.GetFullName(),
				Avatar:   bid.Bidder.Avatar,
			},
		})
	}

	return response
}
