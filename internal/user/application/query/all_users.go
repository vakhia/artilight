package query

import (
	"github.com/vakhia/artilight/internal/user/application/dto"
	"github.com/vakhia/artilight/internal/user/domain/aggregate"
)

type AllUsersHandler struct {
	readModel AllUsersReadModel
}

func NewAllUsersHandler(readModel AllUsersReadModel) AllUsersHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllUsersHandler{readModel: readModel}
}

type AllUsersReadModel interface {
	GetAllUsers() ([]aggregate.User, error)
}

func (h AllUsersHandler) Handle() ([]dto.UserResponse, error) {
	users, err := h.readModel.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponse := mapArtToArtResponse(user)
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func mapArtToArtResponse(user aggregate.User) dto.UserResponse {
	return dto.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Cover:     user.Cover,
		Bio:       user.Bio,
		Currency:  user.Currency,
	}
}
