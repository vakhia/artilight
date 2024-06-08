package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	Id    uuid.UUID
	Slug  string
	Title string
}

func (i Item) TableName() string {
	return "arts"
}
