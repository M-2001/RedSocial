package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/M-2001/RedSocial/bd"
)

/*MostrarBanner servira para mostrar el banner en pantalla*/
func MostrarBanner(w http.ResponseWriter, r *http.Request) {
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
	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Banner no encontrado", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar banner", http.StatusBadRequest)
	}
}
