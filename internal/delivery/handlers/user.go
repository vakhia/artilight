package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/common/base"
	"github.com/vakhia/artilight/internal/common/dto"
	"github.com/vakhia/artilight/internal/common/messages"
	"github.com/vakhia/artilight/internal/usecases"
	"net/http"
)

type UserHandler struct {
	userUseCase usecases.IUserUseCase
}

func NewUserHandler(userUseCase usecases.IUserUseCase) UserHandler {
	return UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var userDTO dto.UserRegisterRequest
	err := ctx.ShouldBind(&userDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgInvalidRequest,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	newUser, err := h.userUseCase.CreateUser(userDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgFailedToCreateNewUser,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusCreated, base.CreateSuccessResponse(
		messages.MsgUserCreatedSuccessfully,
		http.StatusCreated, newUser,
	))
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var userLoginDto dto.UserLoginRequest
	err := ctx.ShouldBind(&userLoginDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgUserLoginFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	user, err := h.userUseCase.Login(userLoginDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgUserLoginFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusCreated, base.CreateSuccessResponse(
		messages.MsgLoginSuccess,
		http.StatusCreated, user,
	))
}
