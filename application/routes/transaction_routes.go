package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jotaGGod/withdrawal-system/application/controller"
)

func HanddleTransactionRoutes(app *fiber.App) {
	app.Add(fiber.MethodPost, "/transaction", controller.CreateTransaction)
}
