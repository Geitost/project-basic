package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"www.autoblox.xyz/server/structs"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&structs.Key{})
	db.AutoMigrate(&structs.CheckoutSession{})

	return db
}
