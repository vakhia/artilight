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

type CategoryHandler struct {
	categoryServices services.ICategoryService
}

func NewCategoryHandler(categoryServices services.ICategoryService) CategoryHandler {
	return CategoryHandler{categoryServices: categoryServices}
}

func (h *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := h.categoryServices.GetAllCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToGetCategories,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgGetCategoriesSuccessfully,
		http.StatusOK, categories,
	))
}

func (h *CategoryHandler) GetCategoryById(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToParse,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	category, err := h.categoryServices.GetCategoryById(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToGetCategories,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgGetCategoriesSuccessfully,
		http.StatusOK, category,
	))
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var request dto.CategoryCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToCreateCategory,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	category, err := h.categoryServices.CreateCategory(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToCreateCategory,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusCreated, base.CreateSuccessResponse(
		messages.MsgCreateCategorySuccessfully,
		http.StatusCreated, category,
	))
}

func (h *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToParse,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	var request dto.CategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToUpdateCategory,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	request.Id = categoryId
	category, err := h.categoryServices.UpdateCategory(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToUpdateCategory,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgUpdateCategorySuccessfully,
		http.StatusOK, category,
	))
}

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToParse,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	err = h.categoryServices.DeleteCategory(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToUpdateCategory,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusNoContent, base.CreateSuccessResponse(
		messages.MsgUpdateCategorySuccessfully,
		http.StatusNoContent, nil,
	))
}
