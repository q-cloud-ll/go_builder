package initialize

import (
	"project/dao/mysql"
	"project/setting"
)

type DsnProvider interface {
	Dsn() string
}

func Gorm() error {
	switch setting.Conf.DbType {
	case "mysql":
		return mysql.GormMysql()
	default:
		return mysql.GormMysql()
	}
}
