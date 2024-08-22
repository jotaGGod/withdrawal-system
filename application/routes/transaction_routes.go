package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jotaGGod/withdrawal-system/application/controller"
)

// HanddleTransactionRoutes sets up the route for handling transaction requests.
//
//	@Summary		Set up transaction routes
//	@Description	Configures the /transaction route for handling POST requests to create a transaction
//	@Tags			transaction
func HanddleTransactionRoutes(app *fiber.App) {
	app.Add(fiber.MethodPost, "/transaction", controller.CreateTransaction)
}
