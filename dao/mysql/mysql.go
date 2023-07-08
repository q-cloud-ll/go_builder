package mysql

import (
	"errors"
	"project/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GormMysql() error {
	m := setting.Conf.MySQLConfig
	if m.DbName == "" {
		return errors.New("db name cannot be empty")
	}
	mysqlCfg := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}

	if sdb, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{}); err != nil {
		return err
	} else {
		sdb.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := sdb.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		db = sdb

		return nil
	}
}

//var db *sqlx.DB
//
//// Init 初始化MySQL连接
//func Init(cfg *setting.MySQLConfig) (err error) {
//	// "user:password@tcp(host:port)/dbname"
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
//	db, err = sqlx.Connect("mysql", dsn)
//	if err != nil {
//		return
//	}
//	db.SetMaxOpenConns(cfg.MaxOpenConns)
//	db.SetMaxIdleConns(cfg.MaxIdleConns)
//	return
//}
//
//// Close 关闭MySQL连接
//func Close() {
//	_ = db.Close()
//}

// gorm connect database.
