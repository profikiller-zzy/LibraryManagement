package flag

import (
	"LibraryManagement/global"
	"LibraryManagement/model"
)

func MakeMigration() {
	var err error
	// 自定义多对多关系表
	err = global.Db.SetupJoinTable(&model.UserModel{}, "BorrowedBooks", &model.UserBorrowBook{})
	if err != nil {
		global.Log.Warn(err.Error())
	}
	// 对模型自动迁移
	err = global.Db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.UserModel{},
			&model.AdminModel{},
			&model.BookModel{},
			model.UserBorrowBook{})
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Info("数据表迁移成功！")
}
