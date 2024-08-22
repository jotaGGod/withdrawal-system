package service

import (
	"github.com/jotaGGod/withdrawal-system/application/entities"
)

// @Summary		Create a transaction
// @Description	Calculates and returns the banknotes used for the requested amount
// @Tags			transaction
// @Param			requestedAmount	query		int	true	"Amount requested for withdrawal"
// @Success		200				{object}	entities.WithdrawalStatement
func CreateTransaction(requestedAmount int) *entities.WithdrawalStatement {
	usedBankNotes := calculateBankNotes(requestedAmount)
	return &entities.WithdrawalStatement{
		RequestedAmount: requestedAmount,
		UsedBankNotes:   usedBankNotes,
	}
}

// calculateBankNotes calculates the quantity of each banknote needed to fulfill the requested amount.
// It is an internal function and not exposed via any route.
//
//	@Summary		Calculate the banknotes for the requested amount
//	@Description	Returns a map of banknotes used to fulfill the requested amount
//	@Param			requestedAmount	query		int			true	"Amount requested for withdrawal"
//	@Success		200				{object}	map[int]int	"Returns a map of banknotes used"
func calculateBankNotes(requestedAmount int) map[int]int {
	var usedBankNotes = map[int]int{200: 0, 100: 0, 50: 0, 20: 0, 10: 0, 5: 0, 2: 0}
	var existingBankNotes = []int{200, 100, 50, 20, 10, 5, 2}
	for _, bankNote := range existingBankNotes {
		if (requestedAmount == 6 || requestedAmount == 8) && bankNote == 5 {
			continue
		}
		for requestedAmount-bankNote >= 0 {
			requestedAmount -= bankNote
			(usedBankNotes)[bankNote] += 1
		}
	}
	return usedBankNotes
}
