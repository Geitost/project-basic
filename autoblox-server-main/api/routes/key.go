package api

import (
	"fmt"
	"math"
	"math/big"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"www.autoblox.xyz/server/config"
	database "www.autoblox.xyz/server/db"
	"www.autoblox.xyz/server/utils"
)

func KeyRoute(db *gorm.DB, hcaptchaKey string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isPro := false

		ip := ctx.Get("CF-Connecting-IP")
		if ip == "" {
			ip = ctx.Get("X-Forwarded-For")
		}

		ip = Ip2Int(net.ParseIP(ip)).String()

		referer := string(ctx.Context().Referer())
		checkoutSessionId := ctx.Query("checkout_session_id")

		// Check if user has just bought pro
		if checkoutSessionId != "" {
			isCheckoutSuccess := utils.ValidateCheckout(db, checkoutSessionId)
			if isCheckoutSuccess {
				isPro = true
			}
		}

		// Get key
		key, err := database.GetKeyByIp(db, ip)
		if err != nil {
			// If key does not exist create key
			if err == gorm.ErrRecordNotFound {
				newKey, err := database.CreateKey(db, ip, isPro)
				if err != nil {
					return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"success": false,
						"message": "Oops! Something went wrong. Please try again later",
					})
				}
				key = newKey
			} else {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"success": false,
					"message": "Oops! Something went wrong. Please try again later",
				})
			}
		}

		// Give user pro if key already exists but user does not have pro
		if isPro {
			key.Pro = true
			key.ExpiresAt = time.Now().Add(time.Duration(config.ProKeyDuration)).UnixMilli()
			err := database.UpdateKey(db, &key)
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"success": false,
					"message": "Oops! Something went wrong. Please try again later",
				})
			}
		}

		// If user is pro then return key to user without checking anything else
		if key.Pro {
			return ctx.Render("key", fiber.Map{
				"key":       key.Value,
				"pro":       true,
				"expiresIn": math.Floor(time.Until(time.UnixMilli(key.ExpiresAt)).Hours() / 24),
			})
		}

		// If user has completed last checkout then return key to user
		if key.Checkpoint == 4 {
			return ctx.Render("key", fiber.Map{
				"key":       key.Value,
				"expiresIn": math.Floor(time.Until(time.UnixMilli(key.ExpiresAt)).Hours()),
			})
		}

		// If user did not come from linkvertise then reset users checkpoint
		if referer != "https://linkvertise.com/" {
			key.Checkpoint = 0
			database.UpdateKey(db, &key)
		}

		// If user has not waited at least 15 seconds between previous checkpoint then display message
		if time.Since(time.UnixMilli(key.UpdatedAt)).Seconds() < 15 && key.Checkpoint != 0 {
			return ctx.Render("redirect", fiber.Map{
				"checkpoint": key.Checkpoint + 1,
				"location":   config.Linkvertises[key.Checkpoint-1],
				"message":    "Complete The Linkvertise",
			})
		}

		// Once user is on last checkpoint flag that the user has earned the key and reset expires at time
		if key.Checkpoint == 3 {
			key.Checkpoint = 4
			key.ExpiresAt = time.Now().Add(time.Duration(config.KeyDuration)).UnixMilli()

			err := database.UpdateKey(db, &key)
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"success": false,
					"message": "Oops! Something went wrong. Please try again later",
				})
			}

			return ctx.Render("key", fiber.Map{
				"key":       key.Value,
				"expiresIn": math.Round(time.Until(time.UnixMilli(key.ExpiresAt)).Hours()),
			})
		}

		// Render checkpoint
		return ctx.Render("checkpoint", fiber.Map{
			"checkpoint":  key.Checkpoint + 1,
			"hcaptchaKey": hcaptchaKey,
		})
	}
}

func ValidateKey(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		ip := ctx.Get("CF-Connecting-IP")
		if ip == "" {
			ip = ctx.Get("X-Forwarded-For")
		}

		ip = Ip2Int(net.ParseIP(ip)).String()

		fmt.Println(ip)

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
			// If key is pro then return valid
			if key.Pro {
				return ctx.JSON(fiber.Map{
					"success": true,
					"valid":   true,
					"pro":     true,
				})
			}

			// If user has completed all checkpoints
			if key.Checkpoint == 4 {
				return ctx.JSON(fiber.Map{
					"success": true,
					"valid":   true,
					"pro":     false,
				})
			}
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid key. Please go to autoblox.xyz/key to get a key",
		})
	}
}

func Ip2Int(ip net.IP) *big.Int {
	i := big.NewInt(0)
	i.SetBytes(ip)
	return i
}
