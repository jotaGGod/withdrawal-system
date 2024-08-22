package entities

type WithdrawalStatement struct {
	RequestedAmount int         `json:"requestedAmount"`
	UsedBankNotes   map[int]int `json:"usedBankNotes"`
}
