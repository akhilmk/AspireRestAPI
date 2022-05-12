package handlers

import (
	"net/http"
	"strconv"

	"github.com/akhilmk/GoRESTAPI/dbaccess"
	"github.com/akhilmk/GoRESTAPI/model"
	"github.com/akhilmk/GoRESTAPI/service"
	"github.com/akhilmk/GoRESTAPI/util"
)

var repo service.IUserRepo

func initRepo() {
	repo = service.GetUserRepo(dbaccess.GetDBSession())
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	data := util.GetQueryParameterFromRequest(r)
	var offset, limit int

	if _, ok := data["offset"]; ok {
		offset, _ = strconv.Atoi(data["offset"])
	}
	if _, ok := data["limit"]; ok {
		limit, _ = strconv.Atoi(data["limit"])
		if limit == 0 || limit > 100 {
			limit = 100 // set default users return count limit to 100
		}
	}

	if repo == nil {
		initRepo()
	}

	users := repo.GetUsers(offset, limit)
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

	if repo == nil {
		initRepo()
	}
	user := repo.GetUser(id)
	res, _ := util.StructToByte(user)
	util.WriteResponseMessage(w, http.StatusOK, res)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := util.ReadReqBodyAsStruct(r, &user)

	if repo == nil {
		initRepo()
	}
	repo.AddUser(user)

	if err == nil {
		res, _ := util.StructToByte(user)
		util.WriteResponseMessage(w, http.StatusOK, res)
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	var user model.User
	err = util.ReadReqBodyAsStruct(r, &user)
	if err != nil {
		util.WriteResponseMessage(w, http.StatusBadRequest, nil)
		return
	}

	if repo == nil {
		initRepo()
	}
	user = repo.UpdateUser(id, user)
	res, _ := util.StructToByte(user)
	util.WriteResponseMessage(w, http.StatusOK, res)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	if repo == nil {
		initRepo()
	}
	user := repo.DeleteUser(id)
	res, _ := util.StructToByte(user)
	util.WriteResponseMessage(w, http.StatusOK, res)
}
