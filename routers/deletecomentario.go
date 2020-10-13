package routers

import (
	"net/http"

	"github.com/M-2001/RedSocial/bd"
)

/*DeleteComentario servira para eliminar un comentario*/
func DeleteComentario(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviarse el parametro id", http.StatusBadRequest)
		return
	}
	err := bd.DeleteComentario(ID, IDUsuario)
	if err != nil {
		http.Error(w, "No se pudo eliminar el comentario Intente mas tarde"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
