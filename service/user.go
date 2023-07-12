package service

import (
	"context"
	"project/consts"
	"project/repository/db/dao"
	"project/repository/db/model"
	"project/types"
	"sync"
)

type UserSrv struct{}

var (
	UserSrvIns  *UserSrv
	UserSrvOnce sync.Once
)

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})

	return UserSrvIns
}

func (s *UserSrv) UserRegisterSrv(ctx context.Context, req *types.UserRegisterReq) (err error) {
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		return err
	}
	if exist {
		err = consts.UserExistErr
	}
	user := &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model.Active,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		return err
	}

	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		return consts.UserCreateErr
	}

	return
}
