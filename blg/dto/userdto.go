package dto

import (
	"blg/blg/model"
	"blg/types"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func UserDto(dbu *model.User, requ *types.ReqUser) bool {
	// dbu.Id = USERID.Add(1)
	dbu.Avatar = requ.Nickname
	dbu.Create_date = time.Now()
	dbu.Email = requ.Account
	// 生成密码的哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requ.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	dbu.Password = string(hashedPassword)
	dbu.Nickname = requ.Nickname
	dbu.Last_login = time.Now()
	dbu.Account = requ.Account
	return true
}
