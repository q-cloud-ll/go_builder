package app

import (
	"context"
	"project/consts"
)

type key int

var userKey key

type UserInfo struct {
	UID      int64  `json:"uid"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, consts.UserInfoErr
	}

	return user, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}
