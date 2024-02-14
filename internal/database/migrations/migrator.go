package migrations

import (
	"github.com/vakhia/artilight/internal/domain"
	"gorm.io/gorm"
)

func Up(database *gorm.DB) {
	createTables(database)
}

func createTables(database *gorm.DB) {
	err := database.AutoMigrate(&domain.User{})
	if err != nil {
		panic("something went wrong")
	}
}
