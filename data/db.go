package data

import (
	"log"
	"mcsm/env"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func MustConnectToDb() {
    db, err := gorm.Open(sqlite.Open(env.DbPath()), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the database\n", err.Error())
    }

    log.Println("Connected to the database")

    if !env.IsProd() {
        db.Logger = logger.Default.LogMode(logger.Info)
    } else {
        db.Logger = logger.Default.LogMode(logger.Warn)
    }

    log.Println("Running Migrations")
    db.AutoMigrate(&User{})

    Db = db
}
