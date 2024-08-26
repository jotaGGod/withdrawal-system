package application

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupApp() *fiber.App {
	app := fiber.New()
	app.Add(fiber.MethodPost, "/transaction", CreateTransaction)
	return app
}

func TestSuccessTransactionIntegration(t *testing.T) {
	app := setupApp()

	tests := []struct {
		name           string
		requestBody    WithdrawalRequest
		expectedStatus int
		expectedBody   WithdrawalStatement
	}{
		{
			name: "Valid withdrawal request with amount 100",
			requestBody: WithdrawalRequest{
				Amount: 100,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 100,
				UsedBankNotes:   map[int]int{2: 0, 5: 0, 10: 0, 20: 0, 50: 0, 100: 1, 200: 0},
			},
		},
		{
			name: "Valid withdrawal request with amount 2222",
			requestBody: WithdrawalRequest{
				Amount: 2222,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 2222,
				UsedBankNotes:   map[int]int{2: 1, 5: 0, 10: 0, 20: 1, 50: 0, 100: 0, 200: 11},
			},
		},
		{
			name: "Valid withdrawal request with amount 30333",
			requestBody: WithdrawalRequest{
				Amount: 30333,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 30333,
				UsedBankNotes:   map[int]int{2: 4, 5: 1, 10: 0, 20: 1, 50: 0, 100: 1, 200: 151},
			},
		},
		{
			name: "Valid withdrawal request with amount 400451",
			requestBody: WithdrawalRequest{
				Amount: 400451,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 400451,
				UsedBankNotes:   map[int]int{2: 3, 5: 1, 10: 0, 20: 2, 50: 0, 100: 0, 200: 2002},
			},
		},
		{
			name: "Valid withdrawal request with amount 5005556",
			requestBody: WithdrawalRequest{
				Amount: 5005556,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 5005556,
				UsedBankNotes:   map[int]int{2: 3, 5: 0, 10: 0, 20: 0, 50: 1, 100: 1, 200: 25027},
			},
		},
		{
			name: "Valid withdrawal request with amount 60006667",
			requestBody: WithdrawalRequest{
				Amount: 60006667,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 60006667,
				UsedBankNotes:   map[int]int{2: 1, 5: 1, 10: 1, 20: 0, 50: 1, 100: 0, 200: 300033},
			},
		},
		{
			name: "Valid withdrawal request with amount 700007759",
			requestBody: WithdrawalRequest{
				Amount: 700007759,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 700007759,
				UsedBankNotes:   map[int]int{2: 2, 5: 1, 10: 0, 20: 0, 50: 1, 100: 1, 200: 3500038},
			},
		},
		{
			name: "Valid withdrawal request with amount 8000880829",
			requestBody: WithdrawalRequest{
				Amount: 8000880829,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 8000880829,
				UsedBankNotes:   map[int]int{2: 2, 5: 1, 10: 0, 20: 1, 50: 0, 100: 0, 200: 40004404},
			},
		},
		{
			name: "Valid withdrawal request with amount 11",
			requestBody: WithdrawalRequest{
				Amount: 11,
			},
			expectedStatus: http.StatusOK,
			expectedBody: WithdrawalStatement{
				RequestedAmount: 11,
				UsedBankNotes:   map[int]int{2: 3, 5: 1, 10: 0, 20: 0, 50: 0, 100: 0, 200: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, -1)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			var responseBody WithdrawalStatement
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, responseBody)
		})
	}
}

func TestTransactionIntegrationError(t *testing.T) {
	app := setupApp()

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedError  map[string]interface{}
	}{
		{
			name: "Invalid withdrawal request with invalid field",
			requestBody: map[string]interface{}{
				"amount": "100",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError: map[string]interface{}{
				"status":  "error",
				"message": "Invalid request body",
			},
		},
		{
			name: "Invalid withdrawal request with wrong field name",
			requestBody: map[string]interface{}{
				"amont": "100",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError: map[string]interface{}{
				"status":  "error",
				"message": "Key: 'WithdrawalRequest.Amount' Error:Field validation for 'Amount' failed on the 'required' tag",
			},
		},
		{
			name: "Invalid withdrawal request with amount 0",
			requestBody: map[string]interface{}{
				"amount": 0,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  map[string]interface{}{"status": "error", "message": "Key: 'WithdrawalRequest.Amount' Error:Field validation for 'Amount' failed on the 'required' tag"},
		},
		{
			name: "Invalid withdrawal request with amount 1",
			requestBody: map[string]interface{}{
				"amount": 1,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  map[string]interface{}{"status": "error", "message": "Withdrawal with the requested amount is not possible. Available bank notes: 2, 5, 10, 20, 50, 100, 200."},
		},
		{
			name: "Invalid withdrawal request with amount 3",
			requestBody: map[string]interface{}{
				"amount": 3,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  map[string]interface{}{"status": "error", "message": "Withdrawal with the requested amount is not possible. Available bank notes: 2, 5, 10, 20, 50, 100, 200."},
		},
		{
			name: "Invalid withdrawal request with negative amount",
			requestBody: map[string]interface{}{
				"amount": -3,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  map[string]interface{}{"status": "error", "message": "Withdrawal with the requested amount is not possible. Available bank notes: 2, 5, 10, 20, 50, 100, 200."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, -1)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			var responseBody map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedError, responseBody)
		})
	}
}
