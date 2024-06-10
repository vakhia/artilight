package command

import (
	"github.com/vakhia/artilight/internal/common/errors"
	"github.com/vakhia/artilight/internal/user/application/dto"
	"github.com/vakhia/artilight/internal/user/domain/repository"
	"github.com/vakhia/artilight/pkg/helper"
	"github.com/vakhia/artilight/pkg/token"
)

type LoginUserHandler struct {
	userRepository repository.UserRepository
	tokenService   token.IJwtService
}

func NewLoginUserHandler(userRepository repository.UserRepository, tokenService token.IJwtService) LoginUserHandler {
	return LoginUserHandler{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (h *LoginUserHandler) Handle(request dto.LoginUserRequest) (dto.UserResponse, error) {
	user, err := h.userRepository.FindByEmail(request.Email)
	if err != nil {
		return dto.UserResponse{}, &errors.UnauthorizedError{Message: "did not find user with that credentials"}
	}

	if !helper.CheckPasswordHash(request.Password, user.Password) {
		return dto.UserResponse{}, &errors.UnauthorizedError{Message: "invalid email or password"}
	}

	token, err := h.tokenService.GenerateToken(user.Email, user.Id.String())
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		FullName:  user.GetFullName(),
		Email:     user.Email,
		Avatar:    user.Avatar,
		Token:     token,
		Currency:  user.Currency,
		Balance:   user.Balance,
	}, nil
}
