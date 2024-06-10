package entity

import (
	"github.com/google/uuid"
)

type ArtImage struct {
	Id    uuid.UUID
	Url   string
	ArtId uuid.UUID
}

func NewArtImage(url string) ArtImage {
	return ArtImage{
		Id:  uuid.New(),
		Url: url,
	}
}
