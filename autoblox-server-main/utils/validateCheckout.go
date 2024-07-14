package utils

import (
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"gorm.io/gorm"
	database "www.autoblox.xyz/server/db"
)

func ValidateCheckout(db *gorm.DB, checkout_session_id string) bool {
	// Return false if records already exist
	if database.CheckoutSessionExists(db, checkout_session_id) {
		return false
	}

	// Get checkout session
	checkoutSession, err := session.Get(checkout_session_id, &stripe.CheckoutSessionParams{})

	// If there is error return false
	if err != nil {
		return false
	}

	// Check If Order Was Complete
	if checkoutSession.Status == stripe.CheckoutSessionStatusComplete {
		// If there is no error return true and add record into the db
		if err == nil {
			database.CreateCheckoutSession(db, checkout_session_id)
			return true
		}
	}

	return false
}
