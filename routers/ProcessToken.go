package routers

import (
	"errors"
	"strings"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor usado en el endponit*/
var Email string

/*IDUsuario id devuelto del modelo*/
var IDUsuario string

/*TokenProcess funcion principal para verificar*/
func TokenProcess(token string) (*models.Claim, bool, string, error) {
	key := []byte("Programacion - RedSocial")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("El formato es Invalido")
	}
	token = strings.TrimSpace(splitToken[1])

	token1, err := jwt.ParseWithClaims(token, claims, func(token2 *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, find, _ := bd.ChequeoYaExiste(claims.Email)
		if find == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, find, IDUsuario, nil
	}
	if !token1.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}
	return claims, false, string(""), err
}
