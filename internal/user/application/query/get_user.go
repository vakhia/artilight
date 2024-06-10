package query

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/user/application/dto"
	"github.com/vakhia/artilight/internal/user/domain/aggregate"
)

type GetUserHandler struct {
	readModel GetUserReadModel
}

func NewGetUserHandler(readModel GetUserReadModel) GetUserHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetUserHandler{readModel: readModel}
}

type GetUserReadModel interface {
	FindById(id uuid.UUID) (art aggregate.User, err error)
}

func (h GetUserHandler) Handle(id uuid.UUID) (dto.UserResponse, error) {
	user, err := h.readModel.FindById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return mapUserToUserResponse(user), err
}

func mapUserToUserResponse(user aggregate.User) dto.UserResponse {
	return dto.UserResponse{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		FullName:    user.GetFullName(),
		Email:       user.Email,
		Avatar:      user.Avatar,
		Cover:       user.Cover,
		Bio:         user.Bio,
		Position:    user.Position,
		Gender:      user.Gender,
		Currency:    user.Currency,
		Balance:     user.Balance,
		PhoneNumber: user.PhoneNumber,
		Location:    user.Location,
		Address:     user.Address,
	}
}
