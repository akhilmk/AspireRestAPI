package model

import (
	"time"
)

type LoanData struct {
	LoanID  int `json:"loanid,omitempty"`
	Amount  int `json:"amount"`
	Terms   int `json:"terms"`
	RepayID int `json:"repayid,omitempty"`
}
type RepayData struct {
	RepayID     int         `json:"repayid,omitempty"`
	Dates       []time.Time `json:"repays"`
	RepayAmount int
	Status      string
}

const (
	PENDING  = "PENDING"
	APPROVED = "APPROVED"
)
