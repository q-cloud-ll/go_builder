package service

import (
	"project/dao/mysql"
	"project/model"
	"project/utils/snowflake"
)

func SignUpService(u *model.UserParamReq) (err error) {
	if err := mysql.CheckUserExist(u.UserName); err != nil {
		return err
	}

	userId := snowflake.GenID()

	user := &model.User{
		UserId:   userId,
		UserName: u.UserName,
		Password: u.Password,
	}

	return mysql.CreateUser(user)
}
