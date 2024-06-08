package command

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/user/application/dto"
	"github.com/vakhia/artilight/internal/user/domain/repository"
)

type UpdateUserHandler struct {
	userRepository repository.UserRepository
}

func NewUpdateUserHandler(userRepository repository.UserRepository) UpdateUserHandler {
	return UpdateUserHandler{
		userRepository: userRepository,
	}
}

func (h *UpdateUserHandler) Handle(id uuid.UUID, request dto.UpdateUserRequest) error {
	user, err := h.userRepository.FindById(id)
	if err != nil {
		return err
	}

	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}
	if request.LastName != "" {
		user.LastName = request.LastName
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Bio != "" {
		user.Bio = request.Bio
	}
	if request.Position != "" {
		user.Position = request.Position
	}
	if request.Gender != "" {
		user.Gender = request.Gender
	}
	if request.Currency != "" {
		user.Currency = request.Currency
	}
	if request.PhoneNumber != "" {
		user.PhoneNumber = request.PhoneNumber
	}
	if request.Location != "" {
		user.Location = request.Location
	}
	if request.Address != "" {
		user.Address = request.Address
	}

	err = h.userRepository.Save(user)
	return err
}
