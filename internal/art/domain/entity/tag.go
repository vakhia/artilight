package entity

import (
	"github.com/google/uuid"
)

type ArtTag struct {
	ArtId uuid.UUID
	TagId int
}

type Tag struct {
	Id          int
	Title       string
	Description string
}

func NewTag(name, description string) Tag {
	return Tag{
		Title:       name,
		Description: description,
	}
}
