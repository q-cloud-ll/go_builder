package mysql

import (
	"database/sql"
	"project/model"
)

func CheckUserExist(userName string) (err error) {
	sqlStr := `select count(user_id) from sys_user where user_name = ?`
	var count int64
	if err := db.Get(&count, sqlStr, userName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func CreateUser(u *model.User) (err error) {
	sqlStr := `insert into sys_user(user_id,user_name,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, u.UserId, u.UserName, u.Password)
	return
}

func SignIn(user *model.User) (err error) {
	password := user.Password
	sqlStr := `select user_id, student_id, user_name, password from sys_user where user_name = ?`
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}

	if password != user.Password {
		return ErrorInvalidPassword
	}

	return
}
