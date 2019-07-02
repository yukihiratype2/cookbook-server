package userservice

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashAndSaltPassword(clearPassword string) string {
	bytePwd := []byte(clearPassword)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash[:])
}

func verifyPassword(hashedPassword string, passwordToVerify string) error {
	byteHashedPwd := []byte(hashedPassword)
	bytePwdToVerify := []byte(passwordToVerify)
	verifyResult := bcrypt.CompareHashAndPassword(byteHashedPwd, bytePwdToVerify)
	return verifyResult
}
