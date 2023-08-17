package util

import "golang.org/x/crypto/bcrypt"

func PasswordToHash(pwd string) (string, error) {
	hashedPwdBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hashedPwdBytes), err
}

func VertifyPasswordHash(pwdhs, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdhs), []byte(pwd))
	return err == nil
}
