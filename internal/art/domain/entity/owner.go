package entity

import "github.com/google/uuid"

type Owner struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Avatar    string
}

func (o Owner) GetFullName() string {
	return o.FirstName + " " + o.LastName
}

func (Owner) TableName() string {
	return "users"
}
