package middlew

import (
	"net/http"

	"github.com/M-2001/RedSocial/routers"
)

/* ValidacionJWT se encarga de la validacion del token*/
func ValidacionJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.TokenProcess(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token !"+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
