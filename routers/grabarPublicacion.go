package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*GrabarPublicacion sirve para grabar un publicacion en la base de datos*/
func GrabarPublicacion(w http.ResponseWriter, r *http.Request) {

	/*imagen publicacion*/
	file, handler, err := r.FormFile("foto")
	var ext = strings.Split(handler.Filename, ".")[1]
	var file1 string = "uploads/publicaciones/" + IDUsuario + "." + ext
	f, err := os.OpenFile(file1, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error con la imagen"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Ocurrio un error Intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}
	//var msj models.GraboPublicacion
	// tecnologias:= r.URL.Query().Get("tecnologia")
	//err := json.NewDecoder(r.Body).Decode(&msj)

	publicacion := r.FormValue("publicacion")
	code := r.FormValue("code")
	tecnologias := r.FormValue("tecnologias")

	registro := models.GraboPublicacion{
		UserID:           IDUsuario,
		Publicacion:      publicacion,
		Code:             code,
		Tecnologias:      tecnologias,
		Foto:             IDUsuario + "." + ext,
		FechaPublicacion: time.Now(),
	}
	_, status, err := bd.InsertPublicacion(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar la publicacion"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la Publicacion", 400)
		return
	}
	w.Header().Set("Content-type", "ParseMultipartForm/form-data")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registro)
}
