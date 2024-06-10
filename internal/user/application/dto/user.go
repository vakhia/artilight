package dto

import "github.com/google/uuid"

type UserResponse struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	FullName    string    `json:"full_name"`
	Avatar      string    `json:"avatar,omitempty"`
	Cover       string    `json:"cover,omitempty"`
	Bio         string    `json:"bio,omitempty"`
	Position    string    `json:"position,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Email       string    `json:"email"`
	Currency    string    `json:"currency,omitempty"`
	Balance     float64   `json:"balance,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Location    string    `json:"location,omitempty"`
	Address     string    `json:"address,omitempty"`
	Token       string    `json:"token,omitempty"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Bio         string `json:"bio,omitempty"`
	Position    string `json:"position,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Currency    string `json:"currency,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Location    string `json:"location,omitempty"`
	Address     string `json:"address,omitempty"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
