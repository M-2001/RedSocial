package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*UpAvatar permitira actualizar informacion del usuario y subir un avatar*/
func UpAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var ext = strings.Split(handler.Filename, ".")[1]
	var file1 string = "uploads/avatars/" + IDUsuario + "." + ext
	f, err := os.OpenFile(file1, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al actualizar Avatar"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Ocurrio un error Intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.Usuario
	var status bool

	user.Avatar = IDUsuario + "." + ext
	status, err = bd.ModificarRegistro(user, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al guardar avatar en base de datos!"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "applictaion/json")
	w.WriteHeader(http.StatusCreated)

}
