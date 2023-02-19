package main

import (
	"github.com/Hana-ame/moonchan-backend/db"
	"github.com/Hana-ame/moonchan-backend/utils"
	"github.com/Hana-ame/moonchan-backend/webfinger"

	"github.com/gofiber/fiber/v2"
)

func webfingerApp() *fiber.App {
	app := fiber.New()

	app.All("", func(c *fiber.Ctx) error {
		acct := c.Query("resource")

		username, domain := utils.ParseAcct(acct)

		if !webfinger.CheckDomain(domain) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "invalid_resource",
			})
		}

		if !db.CheckUserExist(username) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "no_such_user",
			})
		}

		err := c.Status(fiber.StatusOK).JSON(webfinger.GetResource(username, domain))

		return err
	})

	return app
}
