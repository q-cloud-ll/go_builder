package mysql

import (
	"project/model"

	"gorm.io/gorm"
)

func CheckUserExist(userName string) (err error) {
	var count int64
	err = db.Model(model.User{}).Select("user_id").Where("user_name = ?", userName).Count(&count).Error
	if count > 0 {
		return ErrorUserExist
	}

	return
}

func CreateUser(u *model.User) (err error) {
	err = db.Model(model.User{}).Create(u).Error
	return
}

func SignIn(user *model.User) (err error) {
	var u model.User
	password := user.Password
	err = db.Model(model.User{}).Select("user_id, student_id, user_name, password").Where("user_name = ?", user.UserName).Find(&u).Error
	if err == gorm.ErrRecordNotFound {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if password != u.Password {
		return ErrorInvalidPassword
	}

	return
}
