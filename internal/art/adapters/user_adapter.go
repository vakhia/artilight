package adapters

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"github.com/vakhia/artilight/internal/user/domain/repository"
)

type UserAdapter struct {
	userRepository repository.UserRepository
}

func NewUserAdapter(userRepository repository.UserRepository) *UserAdapter {
	return &UserAdapter{userRepository: userRepository}
}

func (a *UserAdapter) FindOwnerById(userID uuid.UUID) (entity.Owner, error) {
	user, err := a.userRepository.FindById(userID)
	if err != nil {
		return entity.Owner{}, err
	}

	return entity.Owner{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
	}, nil
}
