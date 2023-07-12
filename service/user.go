package service

import (
	"context"
	"project/consts"
	"project/logger"
	"project/repository/db/model"
	"project/service/svc"
	"project/types"
	"project/utils/jwt"
	"project/utils/snowflake"

	"go.uber.org/zap"
)

type UserSrv struct {
	ctx    context.Context
	svcCtx *svc.UserServiceContext
	log    *zap.Logger
}

func NewUserService(ctx context.Context, svcCtx *svc.UserServiceContext) *UserSrv {
	return &UserSrv{
		ctx:    ctx,
		svcCtx: svcCtx,
		log:    logger.Lg,
	}
}

func (l *UserSrv) UserRegisterSrv(req *types.UserRegisterReq) (err error) {
	_, exist, err := l.svcCtx.UserModel.ExistOrNotByUserName(l.ctx, req.UserName)
	if err != nil {
		return err
	}
	if exist {
		err = consts.UserExistErr
	}
	user := &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		UserId:   snowflake.GenID(),
		Status:   model.Active,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		return err
	}

	// 创建用户
	err = l.svcCtx.UserModel.CreateUser(l.ctx, user)
	if err != nil {
		return consts.UserCreateErr
	}

	return
}

func (l *UserSrv) UserLoginSrv(req *types.UserRegisterReq) (resp interface{}, err error) {
	var user *model.User
	user, exist, err := l.svcCtx.UserModel.ExistOrNotByUserName(l.ctx, req.UserName)

	if !exist {
		return nil, consts.UserNotExistErr
	}

	if !user.CheckPassword(req.Password) {
		return nil, consts.UserInvalidPasswordErr
	}
	b := jwt.BaseClaims{
		UID:      user.UserId,
		ID:       user.ID,
		Username: user.UserName,
	}
	accessToken, refreshToken, err := jwt.NewJWT().GenerateToken(b)
	if err != nil {
		return nil, err
	}
	userResp := &types.UserInfoResp{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		CreateAt: user.CreatedAt.Unix(),
	}

	resp = &types.UserTokenData{
		User:         userResp,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}
