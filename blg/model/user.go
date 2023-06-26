package model

import (
	"time"
)

type User struct {
	Id                int64     `gorm:"column:id"`
	Account           string    `gorm:"column:account"`
	Admin             []byte    `gorm:"column:admin"`
	Avatar            string    `gorm:"column:avatar"`
	CreateDate        time.Time `gorm:"column:create_date"`
	Deleted           []byte    `gorm:"column:deleted"`
	Email             string    `gorm:"column:email"`
	LastLogin         time.Time `gorm:"column:last_login"`
	MobilePhoneNumber string    `gorm:"column:mobile_phone_number"`
	Nickname          string    `gorm:"column:nickname"`
	Password          string    `gorm:"column:password"`
	Salt              string    `gorm:"column:salt"`
	Status            string    `gorm:"column:status"`
}

func (u *User) TableName() string {
	return "sys_user"
}
