package database

import (
	"gorm.io/gorm"
	"www.autoblox.xyz/server/structs"
)

func CheckoutSessionExists(db *gorm.DB, checkout_session_id string) bool {
	var checkoutSession structs.CheckoutSession

	result := db.Where(structs.CheckoutSession{Value: checkout_session_id}).First(&checkoutSession)

	return result.Error == nil
}

func CreateCheckoutSession(db *gorm.DB, checkout_session_id string) (structs.CheckoutSession, error) {
	checkoutSession := structs.CheckoutSession{Value: checkout_session_id}

	result := db.Create(&checkoutSession)
	if result.Error != nil {
		return structs.CheckoutSession{}, result.Error
	}

	return checkoutSession, nil
}

func GetCheckoutCount(db *gorm.DB) int64 {
	var checkoutCount int64
	db.Model(&structs.Key{}).Count(&checkoutCount)

	return checkoutCount
}
