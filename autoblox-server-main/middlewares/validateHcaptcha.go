package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"www.autoblox.xyz/server/structs"
)

func ValidateHcaptcha(hcaptchaSecret string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := new(structs.HCaptchaReq)

		if err := ctx.BodyParser(data); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Oops! There was an error while trying to process your request",
			})
		}

		// Validate the hCaptcha token by making a request to the hCaptcha API
		response, err := http.PostForm("https://hcaptcha.com/siteverify", map[string][]string{
			"secret":   {hcaptchaSecret},
			"response": {data.Token},
		})

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Sorry, we couldn't validate the hCaptcha. Please try again",
			})
		}
		defer response.Body.Close()

		// Parse the hCaptcha response using go-json
		var hcaptchaResponse structs.HCaptchaRes
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&hcaptchaResponse); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Oops! There was an error while trying to process your request",
			})
		}

		// Check the "success" field in the decoded response
		if hcaptchaResponse.Success {
			return ctx.Next()
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Oops! It looks like there was an issue with the hCaptcha. Please double-check your input and try again",
			})
		}
	}
}
