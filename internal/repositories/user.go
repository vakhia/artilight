package repositories

import (
	"github.com/vakhia/artilight/internal/domain"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUserById(id int) (domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	DeleteUser(id int) error
	GetAllUsers() ([]domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserById(id int) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id int) error {
	err := r.db.Delete(&domain.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
