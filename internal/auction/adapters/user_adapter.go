package adapters

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/auction/domain/entity"
	"github.com/vakhia/artilight/internal/user/domain/repository"
)

type UserAdapter struct {
	userRepository repository.UserRepository
}

func NewUserAdapter(userRepository repository.UserRepository) *UserAdapter {
	return &UserAdapter{userRepository: userRepository}
}

func (a *UserAdapter) FindBidderById(userID uuid.UUID) (entity.Bidder, error) {
	user, err := a.userRepository.FindById(userID)
	if err != nil {
		return entity.Bidder{}, err
	}

	return entity.Bidder{
		Id:       user.Id,
		FullName: user.GetFullName(),
	}, nil
}
