package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/sqlite3"
)

var storage = sqlite3.New(sqlite3.Config{
	Database: "main.db",
	Table:    "ratelimits",
})

var LightLimiter = limiter.New(limiter.Config{
	Max:        15,
	Expiration: 5 * time.Second,
	KeyGenerator: func(c *fiber.Ctx) string {
		return fmt.Sprintf("%s-light", c.IP())
	},
	Storage: storage,
})

var ModerateLimiter = limiter.New(limiter.Config{
	Max:        8,
	Expiration: 5 * time.Second,
	KeyGenerator: func(c *fiber.Ctx) string {
		return fmt.Sprintf("%s-moderate", c.IP())
	},
	Storage: storage,
})

var StrictLimiter = limiter.New(limiter.Config{
	Max:        3,
	Expiration: 10 * time.Second,
	KeyGenerator: func(c *fiber.Ctx) string {
		return fmt.Sprintf("%s-strict", c.IP())
	},
	Storage: storage,
})
