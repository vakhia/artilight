package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/common/base"
	"github.com/vakhia/artilight/internal/common/dto"
	"github.com/vakhia/artilight/internal/common/messages"
	"github.com/vakhia/artilight/internal/services"
	"net/http"
	"strconv"
)

type ArtHandler struct {
	artServices services.IArtService
}

func NewArtHandler(artServices services.IArtService) ArtHandler {
	return ArtHandler{artServices: artServices}
}

func (h *ArtHandler) GetAllArts(ctx *gin.Context) {
	categories, err := h.artServices.GetAllArts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToGetArts,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgGetArtsSuccessfully,
		http.StatusOK, categories,
	))
}

func (h *ArtHandler) GetArtById(ctx *gin.Context) {
	artId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToParse,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	art, err := h.artServices.GetArtById(artId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToGetArts,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgGetArtsSuccessfully,
		http.StatusOK, art,
	))
}

func (h *ArtHandler) CreateArt(ctx *gin.Context) {
	var art dto.ArtCreateRequest
	err := ctx.ShouldBindJSON(&art)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToCreateArt,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	createdArt, err := h.artServices.CreateArt(art)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToCreateArt,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgCreateArtSuccessfully,
		http.StatusOK, createdArt,
	))
}

func (h *ArtHandler) UpdateArt(ctx *gin.Context) {
	artId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToParse,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	var art dto.ArtUpdateRequest
	err = ctx.ShouldBindJSON(&art)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToUpdateArt,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	art.Id = artId
	updatedArt, err := h.artServices.UpdateArt(art)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToUpdateArt,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgUpdateArtSuccessfully,
		http.StatusOK, updatedArt,
	))
}

func (h *ArtHandler) DeleteArt(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToParse,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	err = h.artServices.DeleteArt(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToUpdateArt,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusNoContent, base.CreateSuccessResponse(
		messages.MsgUpdateArtSuccessfully,
		http.StatusNoContent, nil,
	))
}
