package dao

import (
	"context"
	"project/setting"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"gorm.io/gorm"
)

var _db *gorm.DB

func InitMysql() {
	pathRead := strings.Join([]string{setting.Conf.MySQLConfig.User, ":", setting.Conf.MySQLConfig.Password, "@tcp(", setting.Conf.MySQLConfig.Host, ":", setting.Conf.MySQLConfig.Port, ")/", setting.Conf.MySQLConfig.DbName, "?", setting.Conf.MySQLConfig.Config}, "")
	pathWrite := strings.Join([]string{setting.Conf.MySQLConfig.User, ":", setting.Conf.MySQLConfig.Password, "@tcp(", setting.Conf.MySQLConfig.Host, ":", setting.Conf.MySQLConfig.Port, ")/", setting.Conf.MySQLConfig.DbName, "?", setting.Conf.MySQLConfig.Config}, "")

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	mysqlCfg := mysql.Config{
		DSN:                       pathRead, // DSN data source name
		DefaultStringSize:         191,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据当前 MySQL 版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(setting.Conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(setting.Conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	_ = _db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(pathRead)},
			Replicas: []gorm.Dialector{mysql.Open(pathWrite), mysql.Open(pathWrite)},
			Policy:   dbresolver.RandomPolicy{},
		}))
	_db = _db.Set("gorm:table_options", "charset=utf8mb4")
	err = migrate()
	if err != nil {
		panic(err)
	}
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
