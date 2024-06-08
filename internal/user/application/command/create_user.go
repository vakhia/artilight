package command

import (
	"errors"
	"github.com/vakhia/artilight/internal/user/application/dto"
	"github.com/vakhia/artilight/internal/user/domain/aggregate"
	"github.com/vakhia/artilight/internal/user/domain/repository"
	"github.com/vakhia/artilight/pkg/helper"
)

type CreateUserHandler struct {
	repo repository.UserRepository
}

func NewCreateUserHandler(repo repository.UserRepository) CreateUserHandler {
	return CreateUserHandler{
		repo: repo,
	}
}

func (h *CreateUserHandler) Handle(request dto.CreateUserRequest) error {
	_, err := h.repo.FindByEmail(request.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	//password, err := helper.GenerateRandomString(8)
	//if err != nil {
	//	return err
	//}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user, err := aggregate.NewUser(request.FirstName, request.LastName, request.Email, hashedPassword)
	if err != nil {
		return err
	}
	return h.repo.Save(user)
}
