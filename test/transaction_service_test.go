package test

import (
	"github.com/jotaGGod/withdrawal-system/application/entities"
	"github.com/jotaGGod/withdrawal-system/application/service"
	"reflect"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	tests := []struct {
		name            string
		requestedAmount int
		expectedResult  *entities.WithdrawalStatement
	}{
		{
			name:            "Test with amount 100",
			requestedAmount: 100,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 100,
				UsedBankNotes:   map[int]int{200: 0, 100: 1, 50: 0, 20: 0, 10: 0, 5: 0, 2: 0},
			},
		},
		{
			name:            "Test with amount 222",
			requestedAmount: 222,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 222,
				UsedBankNotes:   map[int]int{200: 1, 100: 0, 50: 0, 20: 1, 10: 0, 5: 0, 2: 1},
			},
		},
		{
			name:            "Test with amount 334",
			requestedAmount: 334,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 334,
				UsedBankNotes:   map[int]int{200: 1, 100: 1, 50: 0, 20: 1, 10: 1, 5: 0, 2: 2},
			},
		},
		{
			name:            "Test with amount 455",
			requestedAmount: 455,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 455,
				UsedBankNotes:   map[int]int{200: 2, 100: 0, 50: 1, 20: 0, 10: 0, 5: 1, 2: 0},
			},
		},
		{
			name:            "Test with amount 556",
			requestedAmount: 556,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 556,
				UsedBankNotes:   map[int]int{200: 2, 100: 1, 50: 1, 20: 0, 10: 0, 5: 0, 2: 3},
			},
		},
		{
			name:            "Test with amount 667",
			requestedAmount: 667,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 667,
				UsedBankNotes:   map[int]int{200: 3, 100: 0, 50: 1, 20: 0, 10: 1, 5: 1, 2: 1},
			},
		},
		{
			name:            "Test with amount 758",
			requestedAmount: 758,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 758,
				UsedBankNotes:   map[int]int{200: 3, 100: 1, 50: 1, 20: 0, 10: 0, 5: 0, 2: 4},
			},
		},
		{
			name:            "Test with amount 829",
			requestedAmount: 829,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 829,
				UsedBankNotes:   map[int]int{200: 4, 100: 0, 50: 0, 20: 1, 10: 0, 5: 1, 2: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.CreateTransaction(tt.requestedAmount)
			if result.RequestedAmount != tt.expectedResult.RequestedAmount {
				t.Errorf("got %v, want %v", result.RequestedAmount, tt.expectedResult.RequestedAmount)
			}
			if !reflect.DeepEqual(result.UsedBankNotes, tt.expectedResult.UsedBankNotes) {
				t.Errorf("got %v, want %v", result.UsedBankNotes, tt.expectedResult.UsedBankNotes)
			}
		})
	}
}
