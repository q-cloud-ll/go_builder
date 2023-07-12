package dao

import "project/repository/db/model"

func migrate() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			model.User{},
		)

	return
}
