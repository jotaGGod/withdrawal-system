package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jotaGGod/withdrawal-system/application/service"
	"log"
)

type WithdrawalRequest struct {
	Amount int `json:"amount" validate:"required,numeric"`
}

var validate = validator.New()

// @Summary		Create a transaction
// @Description	Create a transaction with the requested amount
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			request	body		WithdrawalRequest	true	"Withdrawal Request"
// @Success		200		{object}	entities.WithdrawalStatement
// @Failure		400		{object}	fiber.Map	"Error response with validation or processing issues"
// @Router			/transaction [post]
func CreateTransaction(c *fiber.Ctx) error {
	var request WithdrawalRequest
	if err := c.BodyParser(&request); err != nil {
		log.Println("error parsing request body", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}
	if err := validate.Struct(&request); err != nil {
		log.Println("error validating request body", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	if !isValidAmount(request.Amount) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Withdrawal with the requested amount is not possible. Available bank notes: 2, 5, 10, 20, 50, 100, 200.",
		})
	}

	withdrawalAmount := service.CreateTransaction(request.Amount)
	return c.JSON(withdrawalAmount)
}

func isValidAmount(amount int) bool {
	if amount < 0 {
		return false
	}
	return amount%10 != 3 && amount%10 != 1
}
