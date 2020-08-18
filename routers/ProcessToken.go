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
func TokenProcess(tk string) (*models.Claim, bool, string, error) {
	key := []byte("ProgramacionRedSocial")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(" "), errors.New("el formato es invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
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
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
