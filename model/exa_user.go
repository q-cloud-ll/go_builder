package model

type User struct {
	BaseModel
	UserId   int64  `json:"user" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
}

type UserParamReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
