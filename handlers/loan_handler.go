package handlers

import (
	"net/http"
	"strconv"

	"github.com/akhilmk/AspireRestAPI/model"
	"github.com/akhilmk/AspireRestAPI/service"
	"github.com/akhilmk/AspireRestAPI/util"
)

func GetLoan(w http.ResponseWriter, r *http.Request) {
	data := util.GetQueryParameterFromRequest(r)
	var err error
	id := 0
	if _, ok := data["id"]; ok {
		id, err = strconv.Atoi(data["id"])
		if err != nil {
			util.WriteResponseMessage(w, http.StatusBadRequest, nil)
			return
		}
	}

	loan := service.GetLoan(id)
	res, _ := util.StructToByte(loan)
	util.WriteResponseMessage(w, http.StatusOK, res)
}

func AddLoan(w http.ResponseWriter, r *http.Request) {
	var loanData model.LoanData
	err := util.ReadReqBodyAsStruct(r, &loanData)

	service.AddNewLoan(loanData)

	if err == nil {
		res, _ := util.StructToByte(loanData)
		util.WriteResponseMessage(w, http.StatusOK, res)
	}
}
