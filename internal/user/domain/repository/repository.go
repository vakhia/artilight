package repository

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/user/domain/aggregate"
)

type UserRepository interface {
	FindById(id uuid.UUID) (aggregate.User, error)
	FindByEmail(email string) (aggregate.User, error)
	Save(user aggregate.User) error
}
