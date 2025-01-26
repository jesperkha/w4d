package database

import (
	"log"

	"github.com/jesperkha/w4d/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func New(config config.Config) *Database {
	db, err := gorm.Open(sqlite.Open(config.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		db: db,
	}
}

func (db *Database) Migrate() error {
	return db.db.AutoMigrate(&Recipe{})
}

func (db *Database) Close() error {
	sqlDb, err := db.db.DB()
	if err != nil {
		return err
	}

	return sqlDb.Close()
}
