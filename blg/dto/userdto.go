package dto

import (
	"blg/blg/model"
	"blg/types"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func UserDto(dbu *model.User, requ *types.ReqUser) bool {
	// dbu.Id = USERID.Add(1)
	dbu.Avatar = "https://s1.ax1x.com/2023/06/25/pCNLtdP.png"
	dbu.CreateDate = time.Now()
	dbu.Email = fmt.Sprintf("%s@gmail.com", requ.Account)
	// 生成密码的哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requ.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	dbu.Password = string(hashedPassword)
	dbu.Nickname = requ.NickName
	dbu.LastLogin = time.Now()
	dbu.Account = requ.Account
	dbu.Admin = []byte{0}
	dbu.Deleted = []byte{0}
	dbu.MobilePhoneNumber = fmt.Sprintf("%s123", requ.NickName)
	dbu.Salt = "0"
	dbu.Status = "1"
	return true
}
