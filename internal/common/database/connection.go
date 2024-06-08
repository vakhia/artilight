package database

import (
	"fmt"
	"github.com/vakhia/artilight/internal/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DB, cfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Up(db)
	return db
}

func Close(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = dbSQL.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
