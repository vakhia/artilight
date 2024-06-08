package database

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func Up(database *gorm.DB) {
	createTables()
}

func createTables() {

}
