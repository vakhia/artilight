package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/application"
	"github.com/vakhia/artilight/internal/art/application/dto"
	"github.com/vakhia/artilight/internal/common/dtos"
	"github.com/vakhia/artilight/internal/common/server"
	"io"
	"strconv"
)

type HttpServer struct {
	app application.Application
}

func NewHttpServer(app application.Application) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) GetAllArts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	sortBy := ctx.DefaultQuery("sortBy", "")
	sortOrder := ctx.DefaultQuery("sortOrder", "")

	paginateParams := dtos.PaginationParams{
		PageNumber: page,
		PageSize:   pageSize,
	}

	sortingParams := dtos.SortingParams{
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	arts, err := h.app.Queries.AllArts.Handle(paginateParams, sortingParams)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"data": arts})
}

func (h HttpServer) GetArt(ctx *gin.Context) {
	id := ctx.Param("id")

	artId, err := uuid.Parse(id)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	art, err := h.app.Queries.GetArtById.Handle(artId)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"data": art})
}

func (h HttpServer) GetArtBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	art, err := h.app.Queries.GetArtBySlug.Handle(slug)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"data": art})

}

func (h HttpServer) CreateArt(ctx *gin.Context) {
	var request dto.CreateArtRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	err := h.app.Commands.CreateArtCommand.Handle(request)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "Art created successfully"})
}

func (h HttpServer) UploadArtImage(ctx *gin.Context) {
	artId := ctx.Param("id")

	id, err := uuid.Parse(artId)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	f, uploadedFile, err := ctx.Request.FormFile("file")
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}
	defer f.Close()

	fileData, err := io.ReadAll(f)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	err = h.app.Commands.UploadArtImageCommand.Handle(id, fileData, uploadedFile.Filename)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "Image uploaded successfully"})
}

func (h HttpServer) CreateCategory(ctx *gin.Context) {
	var request dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	err := h.app.Commands.CreateCategoryCommand.Handle(request)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "Category created successfully"})
}

func (h HttpServer) CreateCollection(ctx *gin.Context) {
	var request dto.CreateCollectionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	err := h.app.Commands.CreateCollectionCommand.Handle(request)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{"message": "Collection created successfully"})
}

func (h HttpServer) GetAllCollections(ctx *gin.Context) {
	collections, err := h.app.Queries.AllCollection.Handle()
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"data": collections})
}

func (h HttpServer) GetAllCategories(ctx *gin.Context) {
	categories, err := h.app.Queries.AllCategories.Handle()
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"data": categories})
}
