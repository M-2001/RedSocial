package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/M-2001/RedSocial/bd"
)

/*MostrarAvatar servira para mostrar los avatar en pantalla*/
func MostrarAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviarse un ID", http.StatusBadRequest)
		return
	}
	profile, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}
	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Avatar no encontrado", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar avatar", http.StatusBadRequest)
	}
}
