package connections

import "golang.org/x/crypto/bcrypt"

// Encrypt Realiza un hashe de una cadena
func Encrypt(s string) (string, error) {
	sEncrypt, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	return string(sEncrypt), err
}
