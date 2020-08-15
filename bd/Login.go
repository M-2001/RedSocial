package bd

import (
	"github.com/M-2001/RedSocial/models"
	"golang.org/x/crypto/bcrypt"
)

/*Login se encarga de loguear*/
func Login(email string, password string) (models.Usuario, bool) {
	user, find, _ := ChequeoYaExiste(email)
	if find == false {
		return user, false
	}
	passwordConvert := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordConvert)
	if err != nil {
		return user, false
	}
	return user, true
}
