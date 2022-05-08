package service

import (
	"log"

	"github.com/akhilmk/GoRESTAPI/dbaccess"
	"github.com/akhilmk/GoRESTAPI/model"
)

func AddUser(user model.User) {
	result := dbaccess.GetDBSession().Create(&user)
	if result != nil && result.Error != nil {
		log.Println("Error Addding data")
	}
}

func GetUser(id int) model.User {
	//return db.GetLoan(id)
	return model.User{}
}
