package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

var passwordHashBytes = 16

func ScryptPwd(pwd string, salt string) (string, error) {
	h, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, passwordHashBytes)
	if err != nil {
		return "", errors.New("error: failed to generate password hash")
	}
	return fmt.Sprintf("%x", h), nil
}

// GenerateSalt @Description "生成用户的加密的钥匙|generate salt"
// @return   salt 		string    "生成用户的加密的钥匙"
// @return   err   		error     "错误信息"
func GenerateSalt() (salt string, err error) {
	buf := make([]byte, passwordHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", errors.New("error: failed to generate user's salt")
	}

	return fmt.Sprintf("%x", buf), nil
}
