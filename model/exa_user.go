package model

type User struct {
	BaseModel
	UserId   int64  `json:"user" gorm:"index;comment:用户id"`
	UserName string `json:"user_name" gorm:"index;comment:用户名"`
	Password string `json:"password" gorm:"not null;comment:密码"`
}

type UserParamReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserSignIn struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
