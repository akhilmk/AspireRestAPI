package handlers

import (
	"net/http"
	"strconv"

	"github.com/akhilmk/GoRESTAPI/model"
	"github.com/akhilmk/GoRESTAPI/service"
	"github.com/akhilmk/GoRESTAPI/util"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
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

	loan := service.GetUser(id)
	res, _ := util.StructToByte(loan)
	util.WriteResponseMessage(w, http.StatusOK, res)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var loanData model.User
	err := util.ReadReqBodyAsStruct(r, &loanData)

	service.AddUser(loanData)

	if err == nil {
		res, _ := util.StructToByte(loanData)
		util.WriteResponseMessage(w, http.StatusOK, res)
	}
}
