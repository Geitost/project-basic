package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	database "www.autoblox.xyz/server/db"
)

func StatsRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		keyCount := database.GetKeyCount(db)
		checkoutCount := database.GetCheckoutCount(db)

		return ctx.JSON(fiber.Map{
			"success":       true,
			"keyCount":      keyCount,
			"checkoutCount": checkoutCount,
		})
	}
}
