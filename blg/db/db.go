package db

import (
	"time"

	log "github.com/RRRRomeo/QLog/api"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// tencent mysql:111.231.28.172
var DB *gorm.DB

func Init() {
	dsn := "root:root@tcp(111.231.28.172)/local?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	log.Printf("db:%v\n", db)
	DB = db

	var user User
	db.Raw("SELECT * FROM t_user WHERE id = ?", 3).Scan(&user)
	log.Debugf("user:%v\n", user)
	tempUser := &User{
		Id:              1002,
		Avatar:          "test",
		Create_time:     time.Now(),
		Email:           "test@gmail.com",
		Nickname:        "nick",
		Password:        "pass",
		Type:            byte(3),
		Update_time:     time.Now(),
		Username:        "username",
		Last_login_time: time.Now(),
		Login_province:  "shanghai",
		Login_lat:       1000,
		Login_city:      "shanghai",
		Login_lng:       1000,
	}
	result := db.Create(tempUser)
	if result.Error != nil {
		log.Debugf("create fail:%s\n", result.Error)
		return
	}
	log.Debugf("create success!\n")
}
