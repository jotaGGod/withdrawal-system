package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jotaGGod/withdrawal-system/application/routes"
	"log"
)

func main() {
	app := fiber.New()

	routes.HanddleTransactionRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Println(err.Error())
	}
}
