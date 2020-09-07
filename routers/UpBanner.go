package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*UpBanner permitira actualizar informacion del usuario y subir un Banner*/
func UpBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var ext = strings.Split(handler.Filename, ".")[1]
	var file1 string = "uploads/banners/" + IDUsuario + "." + ext
	f, err := os.OpenFile(file1, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al actualizar Banner"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Ocurrio un error Intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.Usuario
	var status bool

	user.Banner = IDUsuario + "." + ext
	status, err = bd.ModificarRegistro(user, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al guardar banner en base de datos!"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "applictaion/json")
	w.WriteHeader(http.StatusCreated)

}
