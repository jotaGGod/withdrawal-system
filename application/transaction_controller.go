package application

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

type WithdrawalStatement struct {
	RequestedAmount int         `json:"requestedAmount"`
	UsedBankNotes   map[int]int `json:"usedBankNotes"`
}

type WithdrawalRequest struct {
	Amount int `json:"amount" validate:"required,numeric"`
}

func (a WithdrawalRequest) isInvalidAmount() bool {
	return a.Amount < 0 || a.Amount == 1 || a.Amount == 3
}

var validate = validator.New()

// @Summary		Create a transaction
// @Description	Create a transaction with the requested amount
// @Tags			transaction
// @Accept			json
// @Produce		json
// @Param			request	body		WithdrawalRequest	true	"Withdrawal Request"
// @Success		200		{object}	WithdrawalStatement
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
	if request.isInvalidAmount() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Withdrawal with the requested amount is not possible. Available bank notes: 2, 5, 10, 20, 50, 100, 200.",
		})
	}
	usedBankNotes := calculateBankNotes(request.Amount)
	return c.JSON(&WithdrawalStatement{
		RequestedAmount: request.Amount,
		UsedBankNotes:   usedBankNotes,
	})
}

var existingBankNotes = []int{200, 100, 50, 20, 10, 5, 2}

func calculateBankNotes(requestedAmount int) map[int]int {
	var usedBankNotes = map[int]int{200: 0, 100: 0, 50: 0, 20: 0, 10: 0, 5: 0, 2: 0}
	for _, bankNote := range existingBankNotes {
		if (requestedAmount == 6 || requestedAmount == 8) && bankNote == 5 {
			continue
		}
		for requestedAmount-bankNote >= 0 {
			if requestedAmount-bankNote == 1 || requestedAmount-bankNote == 3 {
				break
			}
			requestedAmount -= bankNote
			(usedBankNotes)[bankNote] += 1
		}
	}
	return usedBankNotes
}
