package dto

import (
	"blg/blg/model"
	"blg/types"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func UserDto(dbu *model.User, requ *types.ReqUser) bool {
	// dbu.Id = USERID.Add(1)
	dbu.Avatar = requ.Name
	dbu.Create_time = time.Now()
	dbu.Email = requ.Email
	// 生成密码的哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requ.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	dbu.Password = string(hashedPassword)
	dbu.Nickname = requ.Name
	dbu.Last_login_time = time.Now()
	dbu.Username = requ.Name
	dbu.Update_time = time.Now()
	return true
}
