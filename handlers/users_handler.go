package handlers

import (
	"net/http"
	"strconv"

	"github.com/akhilmk/GoRESTAPI/model"
	"github.com/akhilmk/GoRESTAPI/service"
	"github.com/akhilmk/GoRESTAPI/util"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := service.GetUsers()
	res, _ := util.StructToByte(users)
	util.WriteResponseMessage(w, http.StatusOK, res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	data := util.GetPathParameterFromRequest(r)
	var err error
	id := 0
	if _, ok := data["id"]; ok {
		id, err = strconv.Atoi(data["id"])
		if err != nil {
			util.WriteResponseMessage(w, http.StatusBadRequest, nil)
			return
		}
	}

	user := service.GetUser(id)
	res, _ := util.StructToByte(user)
	util.WriteResponseMessage(w, http.StatusOK, res)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := util.ReadReqBodyAsStruct(r, &user)

	service.AddUser(user)

	if err == nil {
		res, _ := util.StructToByte(user)
		util.WriteResponseMessage(w, http.StatusOK, res)
	}
}
