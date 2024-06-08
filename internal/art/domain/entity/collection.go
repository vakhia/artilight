package entity

import "github.com/google/uuid"

type Collection struct {
	Id          uuid.UUID
	Slug        string
	Title       string
	Description string
	OwnerId     uuid.UUID `json:"-" gorm:"column:author_id"`
	Owner       Owner     `json:"owner,omitempty"`
}

func NewCollection(slug, title, description string, owner Owner) Collection {
	return Collection{
		Id:          uuid.New(),
		Slug:        slug,
		Title:       title,
		Description: description,
		Owner:       owner,
		OwnerId:     owner.Id,
	}
}
