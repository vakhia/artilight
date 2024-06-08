package entity

import "github.com/google/uuid"

type Bidder struct {
	Id       uuid.UUID
	FullName string
	Avatar   string
}

func (b Bidder) TableName() string {
	return "users"
}
