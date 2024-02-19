package services

import (
	"errors"
	"github.com/vakhia/artilight/internal/common/dto"
	"github.com/vakhia/artilight/internal/domain"
	"github.com/vakhia/artilight/internal/repositories"
	"github.com/vakhia/artilight/pkg/helper"
	"github.com/vakhia/artilight/pkg/token"
)

type IUserService interface {
	CreateUser(user dto.UserRegisterRequest) (dto.UserResponse, error)
	Login(user dto.UserLoginRequest) (dto.UserResponse, error)
	GetAllUsers() ([]domain.User, error)
}

type UserService struct {
	userRepo     repositories.IUserRepository
	tokenService token.IJwtService
}

func NewUserService(userRepo repositories.IUserRepository, tokenService token.IJwtService) *UserService {
	return &UserService{userRepo: userRepo, tokenService: tokenService}
}

func (u *UserService) CreateUser(user dto.UserRegisterRequest) (dto.UserResponse, error) {
	_, err := u.userRepo.GetUserByEmail(user.Email)
	if err == nil {
		return dto.UserResponse{}, errors.New("email already exists")
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}

	createdUser := domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}
	createdUser, err = u.userRepo.CreateUser(createdUser)
	if err != nil {
		return dto.UserResponse{}, err
	}

	authToken, err := u.tokenService.GenerateToken(createdUser.Email, createdUser.Id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		Id:    createdUser.Id,
		Name:  createdUser.Name,
		Email: createdUser.Email,
		Token: authToken,
	}, nil
}

func (u *UserService) Login(user dto.UserLoginRequest) (dto.UserResponse, error) {
	userData, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return dto.UserResponse{}, errors.New("invalid credentials")
	}

	if !helper.CheckPasswordHash(user.Password, userData.Password) {
		return dto.UserResponse{}, errors.New("invalid credentials")
	}

	authToken, err := u.tokenService.GenerateToken(userData.Email, userData.Id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		Id:    userData.Id,
		Name:  userData.Name,
		Email: userData.Email,
		Token: authToken,
	}, nil
}

func (u *UserService) GetAllUsers() ([]domain.User, error) {
	return u.userRepo.GetAllUsers()
}
