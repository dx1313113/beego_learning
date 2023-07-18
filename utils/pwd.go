package utils

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

var salt = "@#$%"

func ScryptPwd(pwd string) string {
	dk, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		print("加密出错")
		return err.Error()
	}
	return string(dk)
}

func BcryptPwd(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) //加密处理
	return string(hash)
}
