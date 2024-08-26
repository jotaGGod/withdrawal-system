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
				UsedBankNotes:   map[int]int{2: 0, 5: 0, 10: 0, 20: 0, 50: 0, 100: 1, 200: 0},
			},
		},
		{
			name:            "Test with amount 2222",
			requestedAmount: 2222,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 2222,
				UsedBankNotes:   map[int]int{2: 1, 5: 0, 10: 0, 20: 1, 50: 0, 100: 0, 200: 11},
			},
		},
		{
			name:            "Test with amount 30334",
			requestedAmount: 30334,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 30334,
				UsedBankNotes:   map[int]int{2: 2, 5: 0, 10: 1, 20: 1, 50: 0, 100: 1, 200: 151},
			},
		},
		{
			name:            "Test with amount 400455",
			requestedAmount: 400455,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 400455,
				UsedBankNotes:   map[int]int{2: 0, 5: 1, 10: 0, 20: 0, 50: 1, 100: 0, 200: 2002},
			},
		},
		{
			name:            "Test with amount 5005556",
			requestedAmount: 5005556,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 5005556,
				UsedBankNotes:   map[int]int{2: 3, 5: 0, 10: 0, 20: 0, 50: 1, 100: 1, 200: 25027},
			},
		},
		{
			name:            "Test with amount 60006667",
			requestedAmount: 60006667,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 60006667,
				UsedBankNotes:   map[int]int{2: 1, 5: 1, 10: 1, 20: 0, 50: 1, 100: 0, 200: 300033},
			},
		},
		{
			name:            "Test with amount 700007759",
			requestedAmount: 700007759,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 700007759,
				UsedBankNotes:   map[int]int{2: 2, 5: 1, 10: 0, 20: 0, 50: 1, 100: 1, 200: 3500038},
			},
		},
		{
			name:            "Test with amount 8000880829",
			requestedAmount: 8000880829,
			expectedResult: &entities.WithdrawalStatement{
				RequestedAmount: 8000880829,
				UsedBankNotes:   map[int]int{2: 2, 5: 1, 10: 0, 20: 1, 50: 0, 100: 0, 200: 40004404},
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
