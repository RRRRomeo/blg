package model

import (
	"time"
)

// -- ----------------------------
// DROP TABLE IF EXISTS `sys_user`;
// CREATE TABLE `sys_user` (
//
//	`id` bigint(20) NOT NULL AUTO_INCREMENT,
//	`account` varchar(64) DEFAULT NULL,
//	`admin` bit(1) DEFAULT NULL,
//	`avatar` varchar(255) DEFAULT NULL,
//	`create_date` datetime DEFAULT NULL,
//	`deleted` bit(1) DEFAULT NULL,
//	`email` varchar(128) DEFAULT NULL,
//	`last_login` datetime DEFAULT NULL,
//	`mobile_phone_number` varchar(20) DEFAULT NULL,
//	`nickname` varchar(255) DEFAULT NULL,
//	`password` varchar(64) DEFAULT NULL,
//	`salt` varchar(255) DEFAULT NULL,
//	`status` varchar(255) DEFAULT NULL,
//	PRIMARY KEY (`id`),
//	UNIQUE KEY `UK_awpog86ljqwb89aqa1c5gvdrd` (`account`),
//	UNIQUE KEY `UK_ahtq5ew3v0kt1n7hf1sgp7p8l` (`email`)
//
// ) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
type User struct {
	Id                  int64     `gorm:"column:id"`
	Account             string    `gorm:"column:account"`
	Admin               []byte    `gorm:"column:admin"`
	Avatar              string    `gorm:"column:avatar"`
	Create_date         time.Time `gorm:"column:create_date"`
	Deleted             []byte    `gorm:"cloumn:deleted"`
	Email               string    `gorm:"column:email"`
	Last_login          time.Time `gorm:"column:last_login"`
	Mobile_phone_number string    `gorm:"column:mobile_phone_number"`
	Nickname            string    `gorm:"column:nickname"`
	Password            string    `gorm:"column:password"`
	Salt                string    `gorm:"column:salt"`
	Status              string    `gorm:"column:status"`
}

func (u *User) TableName() string {
	return "sys_user"
}
