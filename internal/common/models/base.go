package models

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	createdAt time.Time
	updatedAt time.Time
}

func (b *Base) GetId() string {
	return b.Id.String()
}

func (b *Base) GetCreatedAt() time.Time {
	return b.createdAt
}

func (b *Base) GetUpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Base) SetId() {
	b.Id = uuid.New()
}

func NewId() uuid.UUID {
	return uuid.New()
}
