package main

import (
	"github.com/Hana-ame/moonchan-backend/webfinger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	webfinger.SetDomain("moonchan.xyz")

	app := fiber.New()

	app.Mount(webfinger.WebFingerPath, webfingerApp())

	app.Listen(":3000")

}
