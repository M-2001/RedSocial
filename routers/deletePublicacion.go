package routers

import (
	"net/http"

	"github.com/M-2001/RedSocial/bd"
)

/*DeletePublicacion servira para eliminar una publicacion*/
func DeletePublicacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviarse el parametro id", http.StatusBadRequest)
		return
	}
	err := bd.DeletePublicacion(ID, IDUsuario)
	if err != nil {
		http.Error(w, "No se pudo eliminar la Publicacion Intente mas tarde"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
