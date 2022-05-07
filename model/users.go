package model

import (
	"time"
)

type User struct {
	UserID   int `json:"userid,omitempty"`
	UserName int `json:"use"`
	Terms    int `json:"terms"`
	RepayID  int `json:"repayid,omitempty"`
}
type RepayData struct {
	RepayID     int         `json:"repayid,omitempty"`
	Dates       []time.Time `json:"repays"`
	RepayAmount int
	Status      string
}
