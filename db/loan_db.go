package db

import (
	"github.com/akhilmk/AspireRestAPI/model"
)

var loanInfo map[int]model.LoanData
var repayInfo map[int]model.RepayData

func init() {
	// init with size 1000 records.
	if loanInfo == nil {
		loanInfo = make(map[int]model.LoanData, 1000)
	}
	if repayInfo == nil {
		repayInfo = make(map[int]model.RepayData, 1000)
	}
}
func AddNewLoan(loanNew model.LoanData, repayNew model.RepayData) {
	loanid := getNewUniqLoanID()
	repayid := getNewUniqRepayID()

	loanNew.LoanID = loanid
	loanNew.RepayID = repayid
	repayNew.RepayID = repayid

	loanInfo[loanid] = loanNew
	repayInfo[repayid] = repayNew
	//loanInfo = append(loanInfo, loanNew)
	//repayInfo = append(repayInfo, repayNew)
}
func GetLoan(id int) model.LoanData {
	if _, ok := loanInfo[id]; ok {
		return loanInfo[id]
	}
	return model.LoanData{}
}

func getNewUniqLoanID() int {
	id := 0
	for _, ln := range loanInfo {
		if ln.LoanID > id {
			id = ln.LoanID
		}
	}
	id++
	return id
}
func getNewUniqRepayID() int {
	id := 0
	for _, ln := range repayInfo {
		if ln.RepayID > id {
			id = ln.RepayID
		}
	}
	id++
	return id
}
