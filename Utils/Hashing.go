package utils

import "golang.org/x/crypto/bcrypt"

// golang.org/x/crypto
func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(result), err

}

func CheckPasswords(hashPwd, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
	return err == nil

}
