package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	database "www.autoblox.xyz/server/db"
)

func ValidateKey(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ip := ctx.IP()

		headers := ctx.GetReqHeaders()

		reqKey, ok := headers["Authorization"]

		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Invalid key. Please go to autoblox.xyz/key to get a key",
			})
		}

		key, err := database.GetKeyByIp(db, ip)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"success": false,
					"message": "Invalid key. Please go to autoblox.xyz/key to get a key",
				})
			} else {
				fmt.Println("Error while validating key, ", err.Error())
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": "Oops! Something went wrong. Please try again later.",
				})
			}
		}

		// Check if key matches
		if reqKey[0] == key.Value {
			// If key is pro then proceed
			if key.Pro {
				return ctx.Next()
			}

			// If user has completed all checkpoints then proceed
			if key.Checkpoint == 4 {
				return ctx.Next()
			}
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid key. Please go to autoblox.xyz/key to get a key",
		})
	}
}
