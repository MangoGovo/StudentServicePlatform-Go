package database

import (
	"StuService-Go/internal/global"
	"StuService-Go/pkg/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	user := global.Config.GetString("mysql.user")
	pass := global.Config.GetString("mysql.pass")
	host := global.Config.GetString("mysql.host")
	port := global.Config.GetString("mysql.port")
	dbname := global.Config.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		utils.Log.Fatal(err)
	}

	err = autoMigrate(db)
	if err != nil {
		utils.Log.Fatal(err)

	}

	utils.Log.Println("mysql数据库连接成功")
	return db

}
