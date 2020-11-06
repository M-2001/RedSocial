package jwt

import (
	"time"

	"github.com/M-2001/RedSocial/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT sirve para generar un nuevo token*/
func GeneroJWT(user models.Usuario) (string, error) {
	myKey := []byte("ProgramacionRedSocial")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"nombre":    user.Nombre,
		"apellidos": user.Apellido,
		"fechaN":    user.FechaN,
		"biografia": user.Biografia,
		"ubicacion": user.Ubicacion,
		"sitioweb":  user.SitioWeb,
		"_id":       user.ID.Hex(),
		"expirate":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenS, err := token.SignedString(myKey)
	if err != nil {
		return tokenS, err
	}
	return tokenS, nil
}
