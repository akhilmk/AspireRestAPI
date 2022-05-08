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
	user := model.User{}
	result := dbaccess.GetDBSession().First(&user, id)
	if result.Error != nil {
		log.Println("Error getting user")
	}
	return user
}
func GetUsers() []model.User {
	users := []model.User{}
	result := dbaccess.GetDBSession().Find(&users)
	if result.Error != nil {
		log.Println("Error getting users")
	}
	return users
}
