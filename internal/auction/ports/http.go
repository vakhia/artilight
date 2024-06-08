package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/application"
	"github.com/vakhia/artilight/internal/auction/application/dto"
	"github.com/vakhia/artilight/internal/common/server"
	"github.com/vakhia/artilight/internal/common/ws"
)

type HttpServer struct {
	app application.Application
}

func NewHttpServer(app application.Application) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) CreateAuction(ctx *gin.Context) {
	var request dto.CreateAuctionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	err := h.app.Commands.CreateAuctionCommand.Handle(request)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "Auction created successfully"})
}

func (h HttpServer) UpdateAuction(ctx *gin.Context) {
	id := ctx.Param("id")
	auctionId, err := uuid.Parse(id)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	var request dto.UpdateAuctionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	err = h.app.Commands.UpdateAuctionCommand.Handle(auctionId, request)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Auction updated successfully"})
}

func (h HttpServer) CreateBid(ctx *gin.Context) {
	var request dto.CreateBidRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	err := h.app.Commands.CreateBidCommand.Handle(request)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}
	// Broadcast new bid
	ws.Broadcast <- ws.Message{
		Event: "new_bid",
		Data:  request,
	}

	ctx.JSON(201, gin.H{"message": "Bid created successfully"})
}

func (h HttpServer) GetAuction(ctx *gin.Context) {
	id := ctx.Param("id")
	auctionId, err := uuid.Parse(id)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	auction, err := h.app.Queries.GetAuctionQuery.Handle(auctionId)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{"data": auction})
}
