package bd

import "golang.org/x/crypto/bcrypt"

/*EncryptPass sirve para encriptar la contrase;a*/
func EncryptPass(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
