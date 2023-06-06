package db

import (
	"time"
)

// DROP TABLE IF EXISTS `t_user`;
// CREATE TABLE `t_user`  (
//
//	`id` bigint NOT NULL,
//	`avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
//	`create_time` datetime(6) NULL DEFAULT NULL,
//	`email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
//	`nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
//	`password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
//	`type` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//	`update_time` datetime(6) NULL DEFAULT NULL,
//	`username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
//	`last_login_time` datetime NULL DEFAULT NULL,
//	`login_province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//	`login_city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
//	`login_lat` double NULL DEFAULT NULL,
//	`login_lng` double NULL DEFAULT NULL,
//	PRIMARY KEY (`id`) USING BTREE
//
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;
type User struct {
	Id              int64     `gorm:"column:id"`
	Avatar          string    `gorm:"column:avatar"`
	Create_time     time.Time `gorm:"column:create_time"`
	Email           string    `gorm:"column:email"`
	Nickname        string    `gorm:"column:nickname"`
	Password        string    `gorm:"column:password"`
	Type            byte      `gorm:"column:type"`
	Update_time     time.Time `gorm:"column:update_time"`
	Username        string    `gorm:"column:username"`
	Last_login_time time.Time `gorm:"column:last_login_time"`
	Login_province  string    `gorm:"column:login_province"`
	Login_city      string    `gorm:"column:login_city"`
	Login_lat       float64   `gorm:"column:login_lat"`
	Login_lng       float64   `gorm:"column:login_lng"`
}

func (u *User) TableName() string {
	return "t_user"
}
