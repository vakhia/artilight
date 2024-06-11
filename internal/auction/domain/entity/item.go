package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	Id    uuid.UUID
	Slug  string
	Title string
	Price float64
}

func (i Item) TableName() string {
	return "arts"
}
