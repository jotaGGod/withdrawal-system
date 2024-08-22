package service

import (
	"github.com/jotaGGod/withdrawal-system/application/entities"
)

func CreateTransaction(requestedAmount int) *entities.WithdrawalStatement {
	usedBankNotes := calculateBankNotes(requestedAmount)
	return &entities.WithdrawalStatement{
		RequestedAmount: requestedAmount,
		UsedBankNotes:   usedBankNotes,
	}
}

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
