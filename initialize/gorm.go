package initialize

import (
	"os"
	"project/dao/mysql"
	"project/global"
	"project/model"
	"project/setting"

	"go.uber.org/zap"
)

type DsnProvider interface {
	Dsn() string
}

func Gorm() (err error) {
	switch setting.Conf.DbType {
	case "mysql":
		return mysql.GormMysql()
	default:
		return mysql.GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.GB_MDB
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		global.GB_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GB_LOG.Info("register table success")
}
