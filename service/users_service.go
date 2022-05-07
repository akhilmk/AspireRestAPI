package service

import (
	"log"
	"time"

	"github.com/akhilmk/GoRESTAPI/dbaccess"
	"github.com/akhilmk/GoRESTAPI/dbmodel"
	"github.com/akhilmk/GoRESTAPI/model"
)

func AddUser(loanData model.User) {

	var repayDays []time.Time
	for i := 1; i <= loanData.Terms; i++ {
		repayDays = append(repayDays, time.Now().AddDate(0, 0, i*7))
	}

	//repayAmount := loanData.Amount / loanData.Terms

	//loanInfo := model.LoanData{Amount: loanData.Amount, Terms: loanData.Terms}
	//repayInfo := model.RepayData{RepayAmount: repayAmount, Dates: repayDays, Status: model.PENDING}

	db := dbaccess.GetDBSession()
	result := db.Create(dbmodel.User{})
	if result != nil && result.Error != nil {
		log.Println("Error Addding data")
	}
}

func GetUser(id int) model.User {
	//return db.GetLoan(id)
	return model.User{}
}
