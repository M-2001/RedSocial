package routers

import (
	"io"
	"net/http"
	"os"
	"strings"
	//"github.com/M-2001/RedSocial/bd"
)

/*UpImagePublicacion permite subir foto*/
func UpImagePublicacion(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("foto")
	var ext = strings.Split(handler.Filename, ".")[1]
	var file1 string = "uploads/publicaciones/" + IDUsuario + "." + ext
	f, err := os.OpenFile(file1, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al actualizar publicacion"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Ocurrio un error Intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}
	//var pub models.GraboPublicacion
	var status bool

	//pub.Foto = IDUsuario + "." + ext
	//status, err = bd.ModificarPublicacion(pub, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al guardar imagen de publicacion en base de datos!"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
