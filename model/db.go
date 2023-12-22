package model

import (
	"log"

	"github.com/davidmacdonald11/mcsm/cmd/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

type IdType = uint64

func MustConnectToDb() {
	db, err := gorm.Open(sqlite.Open(env.DbPath()), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database\n", err.Error())
	}

	log.Println("Connected to the database")

	if env.IsProd() {
		db.Logger = logger.Default.LogMode(logger.Warn)
	} else {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	log.Println("Running migrations")
	db.AutoMigrate(&InviteCode{}, &User{})

	Db = db
}
