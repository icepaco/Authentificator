package main

import (
	"golang.org/x/crypto/bcrypt"
)

func main() {

}
func savePassword() ([]byte, error) {
	var pswd = []byte("god")
	return bcrypt.GenerateFromPassword(pswd, 64)
}
