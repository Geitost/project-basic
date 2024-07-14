package utils

import (
	"log"
	"time"

	"gorm.io/gorm"
	"www.autoblox.xyz/server/structs"
)

func DeleteExpiredKeys(db *gorm.DB) {
	expiredKeys := []structs.Key{}
	db.Where("expires_at <= ?", time.Now().UnixMilli()).Find(&expiredKeys)

	// Delete the expired key from the database
	for _, key := range expiredKeys {
		db.Unscoped().Delete(&key)
	}

	log.Println("Deleted Expired Keys")
}
