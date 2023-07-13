package dao

import (
	"context"
	"project/repository/db/model"

	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		CreateUser(ctx context.Context, user *model.User) error
		ExistOrNotByUserName(ctx context.Context, userName string) (user *model.User, exist bool, err error)
	}

	customUserModel struct {
		*gorm.DB
	}
)

func NewUserModel() UserModel {
	return &customUserModel{
		DB: NewDBClient(),
	}
}

func (m *customUserModel) CreateUser(ctx context.Context, user *model.User) error {
	return m.DB.WithContext(ctx).Model(&model.User{}).Create(&user).Error
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (m *customUserModel) ExistOrNotByUserName(ctx context.Context, userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = m.DB.WithContext(ctx).Model(&model.User{}).Where("user_name = ?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = m.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}
