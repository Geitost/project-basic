package api

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/stripe/stripe-go/v75"
	"go.uber.org/zap"
	"gorm.io/gorm"
	api "www.autoblox.xyz/server/api/routes"
	"www.autoblox.xyz/server/config"
	"www.autoblox.xyz/server/middlewares"
)

func Start(db *gorm.DB, stripeKey string, hcaptchaKey string, hcaptchaSecret string) {
	// Set stripe key
	stripe.Key = stripeKey

	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		AppName:               "Autoblox",
		DisableStartupMessage: true,

		// Faster JSON
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		// View engine
		Views: engine,

		EnableIPValidation: true,
	})

	// CORS
	app.Use(cors.New())

	// LOGGER
	logger, _ := zap.NewProduction()
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	// SECURITY
	app.Use(helmet.New())

	// Route for static files
	app.Static("/", "./public")

	// Stats
	app.Get("/stats", api.StatsRoute(db))

	// Get key
	app.Get("/key/*", middlewares.LightLimiter, api.KeyRoute(db, hcaptchaKey))

	// Validate checkpoint
	app.Post("/validate-checkpoint", middlewares.LightLimiter, middlewares.ValidateHcaptcha(hcaptchaSecret), api.ValidateCheckoutRoute(db))

	// Validate Key
	app.Post("/validate-key", middlewares.ModerateLimiter, api.ValidateKey(db))

	// Bloxburg
	app.Post("/bloxburg/cashier", middlewares.StrictLimiter, middlewares.ValidateKey(db), api.BloxburgCashierRoute())

	// Other
	app.Get("/pro", func(c *fiber.Ctx) error {
		return c.Redirect(config.ProUrl)
	})
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Redirect(config.DownloadUrl)
	})
	app.Get("/discord", func(c *fiber.Ctx) error {
		return c.Redirect(config.DiscordUrl)
	})

	log.Printf("Server Available At http://localhost:%v", config.Port)

	app.Listen(fmt.Sprintf("127.0.0.1:%v", config.Port))
}
