package entity

import "github.com/google/uuid"

type Bidder struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Avatar    string
}

func (b Bidder) TableName() string {
	return "users"
}

func (b Bidder) GetFullName() string {
	return b.FirstName + " " + b.LastName
}
