package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/M-2001/RedSocial/bd"
)

/*ReadPublicacionesSeguidores permitira leer las publicaciones de los seguidores*/
func ReadPublicacionesSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina mayor a cero", http.StatusBadRequest)
		return
	}
	resp, correct := bd.ReadPublicacionesSeguidores(IDUsuario, page)
	if correct == false {
		http.Error(w, "Error al leer publicaciones", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
