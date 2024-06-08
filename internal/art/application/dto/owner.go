package dto

import "github.com/google/uuid"

type OwnerResponse struct {
	Id       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Avatar   string    `json:"avatar"`
}
