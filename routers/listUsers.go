package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/M-2001/RedSocial/bd"
)

/*ListUsers se encargara de devolder la lista de ususarios */
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("tipo")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pageTemp)

	result, status := bd.ReadAllUsers(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al intentar leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
