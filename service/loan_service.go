package service

import (
	"log"
	"time"

	"github.com/akhilmk/AspireRestAPI/dbaccess"
	"github.com/akhilmk/AspireRestAPI/model"
)

func AddNewLoan(loanData model.LoanData) {

	var repayDays []time.Time
	for i := 1; i <= loanData.Terms; i++ {
		repayDays = append(repayDays, time.Now().AddDate(0, 0, i*7))
	}

	//repayAmount := loanData.Amount / loanData.Terms

	//loanInfo := model.LoanData{Amount: loanData.Amount, Terms: loanData.Terms}
	//repayInfo := model.RepayData{RepayAmount: repayAmount, Dates: repayDays, Status: model.PENDING}

	db := dbaccess.GetDBSession()
	result := db.Create(dbaccess.Loan{})
	if result != nil && result.Error != nil {
		log.Println("Error Addding data")
	}
}

func GetLoan(id int) model.LoanData {
	//return db.GetLoan(id)
	return model.LoanData{}
}
