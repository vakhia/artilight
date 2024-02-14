package domain

import (
	"time"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
