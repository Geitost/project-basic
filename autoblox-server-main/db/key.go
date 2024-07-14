package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"www.autoblox.xyz/server/config"
	"www.autoblox.xyz/server/structs"
)

func ValidateKey(db *gorm.DB, value string) bool {
	var key structs.Key

	result := db.Where("value = ?", value).First(&key)

	// If no error return true (it will throw an error if key is not found in db)
	return result.Error == nil
}

func ValidateProKey(db *gorm.DB, value string) bool {
	var key structs.Key

	result := db.Where("value = ? AND pro = true", value).First(&key)

	// If no error return true (it will throw an error if key is not found in db)
	return result.Error == nil
}

func ValidateKeyByIp(db *gorm.DB, ip string) bool {
	var key structs.Key

	result := db.Where("ip = ?", ip).First(&key)

	// If no error return true (it will throw an error if key is not found in db)
	return result.Error == nil
}

func GetKeyByIp(db *gorm.DB, ip string) (structs.Key, error) {
	var key structs.Key

	result := db.Where("ip = ?", ip).First(&key)

	return key, result.Error
}

func CreateKey(db *gorm.DB, ip string, pro bool) (structs.Key, error) {
	expiresAt := config.KeyDuration

	// If user is pro then set expiresAt to pro key duration
	if pro {
		expiresAt = config.ProKeyDuration
	}

	key := structs.Key{Ip: ip, Value: uuid.NewString(), Pro: pro, ExpiresAt: time.Now().Add(time.Duration(expiresAt)).UnixMilli()}

	result := db.Create(&key)
	if result.Error != nil {
		return structs.Key{}, result.Error
	}

	return key, nil
}

func UpdateKey(db *gorm.DB, key *structs.Key) error {
	result := db.Save(key)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetKeyCount(db *gorm.DB) int64 {
	var keyCount int64
	db.Model(&structs.Key{}).Count(&keyCount)

	return keyCount
}
