package pwdutil

import (
	"groove-app/pkg/crypto"
	"groove-app/pkg/strings"
)

// SaltHashPwd generates a salt and the salt-hashed password.
// The salt-hashed generation algorithm is:
//
//	saltHashPassword = md5(md5(password) + salt)
func SaltHashPwd(password string, saltLen int) (saltHashPassword, salt string) {
	salt = strings.RandLetterNum(saltLen)
	saltHashPassword = crypto.Md5(crypto.Md5(password) + salt)
	return
}

// VerifySaltHashPwd verifies whether saltHashPassword was generated from password and salt.
func VerifySaltHashPwd(password, salt, saltHashPassword string) bool {
	return crypto.Md5(crypto.Md5(password)+salt) == saltHashPassword
}
