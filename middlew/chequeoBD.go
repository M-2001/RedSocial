package middlew

import (
	"net/http"
	"github.com/M-2001/RedSocial/bd"
)

/*ChequeoBD chequa la conexion a la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Error de conexion", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
