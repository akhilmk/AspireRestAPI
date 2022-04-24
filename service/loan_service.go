package service

import (
	"time"

	"github.com/akhilmk/AspireRestAPI/db"
	"github.com/akhilmk/AspireRestAPI/model"
)

func AddNewLoan(loanData model.LoanData) {

	var repayDays []time.Time
	for i := 1; i <= loanData.Terms; i++ {
		repayDays = append(repayDays, time.Now().AddDate(0, 0, i*7))
	}

	repayAmount := loanData.Amount / loanData.Terms

	loanInfo := model.LoanData{Amount: loanData.Amount, Terms: loanData.Terms}
	repayInfo := model.RepayData{RepayAmount: repayAmount, Dates: repayDays, Status: model.PENDING}

	db.AddNewLoan(loanInfo, repayInfo)

}

func GetLoan(id int) model.LoanData {
	return db.GetLoan(id)
}
