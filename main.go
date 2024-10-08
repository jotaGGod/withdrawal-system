package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jotaGGod/withdrawal-system/application"
	_ "github.com/jotaGGod/withdrawal-system/docs"
	"log"
)

// @title			withdrawal-system API
// @version		1.0
// @description	This is an API for a withdrawal system.
// @termsOfService	http://swagger.io/terms/
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Add(fiber.MethodPost, "/transaction", application.CreateTransaction)
	err := app.Listen(":3000")
	if err != nil {
		log.Println(err.Error())
	}
}
