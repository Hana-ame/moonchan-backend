package main

import (
	"log"

	"github.com/Hana-ame/moonchan-backend/webfinger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	webfinger.SetDomain("n.tsukishima.top")

	app := fiber.New()

	app.Mount(webfinger.WebFingerPath, webfingerApp())
	app.Use("/users/", CheckHttpsig)

	log.Fatal(app.Listen("127.0.1.1:3000"))

}
