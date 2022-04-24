package dbaccess

type Loan struct {
	LoanID  int `gorm:"primarykey"`
	Amount  int
	Term    int
	RepayID int
}
