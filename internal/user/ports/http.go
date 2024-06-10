package ports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/common/server"
	"github.com/vakhia/artilight/internal/user/application"
	"github.com/vakhia/artilight/internal/user/application/dto"
	"net/http"
)

type HttpServer struct {
	app application.Application
}

func NewHttpServer(app application.Application) *HttpServer {
	return &HttpServer{app: app}
}

func (h *HttpServer) CreateUser(ctx *gin.Context) {
	var request dto.CreateUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	if err := h.app.Commands.CreateUser.Handle(request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *HttpServer) LoginUser(ctx *gin.Context) {
	var request dto.LoginUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	data, err := h.app.Commands.LoginUser.Handle(request)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *HttpServer) UploadAvatar(ctx *gin.Context) {
	if err := h.app.Commands.UploadAvatar.Handle(ctx); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Avatar uploaded successfully"})
}

func (h *HttpServer) UploadCover(ctx *gin.Context) {
	if err := h.app.Commands.UploadCover.Handle(ctx); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cover uploaded successfully"})
}

func (h *HttpServer) GetMe(ctx *gin.Context) {
	contextId, exists := ctx.Get("userId")
	if !exists {
		server.RespondWithError(ctx, fmt.Errorf("userId not found in context"))
		return
	}

	contextIdStr, ok := contextId.(string)
	if !ok {
		server.RespondWithError(ctx, fmt.Errorf("userId in context is not a string"))
		return
	}

	contextUUID, err := uuid.Parse(contextIdStr)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	data, err := h.app.Queries.GetUser.Handle(contextUUID)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *HttpServer) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	data, err := h.app.Queries.GetUser.Handle(userId)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *HttpServer) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	var request dto.UpdateUserRequest
	if err := ctx.BindJSON(&request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	if err := h.app.Commands.UpdateUser.Handle(userId, request); err != nil {
		server.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
