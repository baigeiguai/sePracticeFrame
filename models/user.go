package models

import (
	"errors"
	"sePracticeFrame/dao"
	"sePracticeFrame/types"
	"sePracticeFrame/utils"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

/*

	register the user
*/
func RegisterUser(user *User) (state types.RegisterMessage) {
	if ExistUser(user) {
		return types.RegisterNameExisted
	}
	user.Username = utils.Convert(user.Username)
	err := dao.DB.Create(user).Error
	if err != nil {
		return types.RegisterError
	}
	return types.RegisterSuccess
}

// true if the username has existed
// else false
func ExistUser(user *User) bool {
	err := dao.DB.Select("username").Where("username = ?", user.Username).First(&User{})
	return !errors.Is(err.Error, gorm.ErrRecordNotFound)
}

func LoginUser(user *User) (state types.LoginMessage) {
	var t User
	err := dao.DB.Where("username = ?", user.Username).First(&t)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return types.LoginNameNotExist
	} else if err.Error != nil {
		return types.LoginError
	}
	if utils.Convert(user.Password) != t.Password {
		return types.LoginPasswordNotCorrect
	}
	return types.LoginSuccess
}
