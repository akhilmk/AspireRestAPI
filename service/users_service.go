package service

import (
	"log"

	"github.com/akhilmk/GoRESTAPI/dbaccess"
	"github.com/akhilmk/GoRESTAPI/model"
	"gorm.io/gorm"
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
func GetUsers(offset, limit int) []model.User {
	users := []model.User{}
	var result *gorm.DB
	if offset == 0 {
		result = dbaccess.GetDBSession().Limit(limit).Order("user_id desc").Find(&users) // get recent users if offset is zero
	} else {
		result = dbaccess.GetDBSession().Limit(limit).Offset(offset).Find(&users)
	}
	if result.Error != nil {
		log.Println("Error getting users")
	}
	return users
}

func UpdateUser(id int, user model.User) model.User {
	user.UserID = id
	result := dbaccess.GetDBSession().Save(&user)
	if result.Error != nil {
		log.Println("Error updating user")
	}
	return user
}
