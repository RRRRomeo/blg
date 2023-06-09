package db

import (
	"blg/tools/cnf"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// tencent mysql:111.231.28.172
var Global *gorm.DB

func Init() {
	dbcnf := cnf.GlobalCnf.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbcnf.User, dbcnf.Pass, dbcnf.Host, dbcnf.Port, dbcnf.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// log.Printf("db:%v\n", db)
	Global = db
}

func GetDB() *gorm.DB {
	return Global
}
