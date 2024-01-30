package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateConnection(entity any) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(entity)

	return db, nil
}
