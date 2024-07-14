package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"www.autoblox.xyz/server/config"
	database "www.autoblox.xyz/server/db"
)

func ValidateCheckoutRoute(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ip := ctx.IP()

		key, err := database.GetKeyByIp(db, ip)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
			})
		}

		// Get The Checkpoint Value Before Updating
		var checkpoint = key.Checkpoint

		// Update Checkpoint!
		res := db.Model(&key).Update("Checkpoint", key.Checkpoint+1)
		if res.Error != nil {
			fmt.Println("Error While Updating Checkpoint, ", res.Error.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
			})
		}

		return ctx.JSON(fiber.Map{
			"success":  true,
			"location": config.Linkvertises[checkpoint],
		})
	}
}
