package service

import (
	"log"

	"github.com/akhilmk/GoRESTAPI/model"
	"gorm.io/gorm"
)

type IUserRepo interface {
	AddUser(user model.User) error
	GetUser(id int) model.User
	GetUsers(offset, limit int) []model.User
	UpdateUser(id int, user model.User) model.User
	DeleteUser(id int) model.User
}

type userRepo struct {
	DB *gorm.DB
}

func GetUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{
		DB: db,
	}
}

func (u *userRepo) AddUser(user model.User) error {
	result := u.DB.Create(&user)
	if result != nil && result.Error != nil {
		log.Println("Error Addding data")
		return result.Error

	}
	return nil
}

func (u *userRepo) GetUser(id int) model.User {
	user := model.User{}
	result := u.DB.First(&user, id)
	if result.Error != nil {
		log.Println("Error getting user")
	}
	return user
}
func (u *userRepo) GetUsers(offset, limit int) []model.User {
	users := []model.User{}
	var result *gorm.DB
	if offset == 0 {
		result = u.DB.Limit(limit).Order("user_id desc").Find(&users) // get recent users if offset is zero
	} else {
		result = u.DB.Limit(limit).Offset(offset).Find(&users)
	}
	if result.Error != nil {
		log.Println("Error getting users")
	}
	return users
}

func (u *userRepo) UpdateUser(id int, user model.User) model.User {
	user.UserID = id
	result := u.DB.Save(&user)
	if result.Error != nil {
		log.Println("Error updating user")
	}
	return user
}
func (u *userRepo) DeleteUser(id int) model.User {
	user := model.User{}
	result := u.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		log.Println("Error deleting user")
	}
	return user
}
