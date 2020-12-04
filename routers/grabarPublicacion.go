package routers

import (
	"encoding/json"
	"fmt"
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

	const layout = "01-02-2006_15040500"
	t := time.Now()

	redSocial := "GK-"

	/*imagen publicacion*/
	file, handler, err := r.FormFile("foto")
	if err != nil {
		handler.Filename = ".jpg"
		fmt.Fprintf(w, "%s\n", err)
		fmt.Println(err)
		return
	}
	defer file.Close()

	var ext = strings.Split(handler.Filename, ".")[1]
	var file1 string = "uploads/publicaciones/" + redSocial + t.Format(layout) + IDUsuario + "." + ext
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

	// RedSocial := time.

	// fechS, err := time.Parse(time.ANSIC, RedSocial.String())

	// var times = fechS

	publicacion := r.FormValue("publicacion")
	code := r.FormValue("code")
	tecnologias := r.FormValue("tecnologias")

	registro := models.GraboPublicacion{
		UserID:           IDUsuario,
		Publicacion:      publicacion,
		Code:             code,
		Tecnologias:      tecnologias,
		Foto:             redSocial + t.Format(layout) + IDUsuario + "." + ext,
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

/*GrabarPublicacionJSON sirve para grabar un publicacion en la base de datos*/
func GrabarPublicacionJSON(w http.ResponseWriter, r *http.Request) {

	var msj models.GraboPublicacion
	// tecnologias:= r.URL.Query().Get("tecnologia")
	err := json.NewDecoder(r.Body).Decode(&msj)

	registro := models.GraboPublicacion{
		UserID:      IDUsuario,
		Publicacion: msj.Publicacion,
		Code:        msj.Code,
		Tecnologias: msj.Tecnologias,
		// Foto:             redSocial + t.Format(layout) + IDUsuario + "." + ext,
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

/*MostrarFotoPublicacion servira para mostrar las fotos de las publicaciones*/
func MostrarFotoPublicacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviarse un ID", http.StatusBadRequest)
		return
	}
	pub, err := bd.BuscoPublicacion(ID)
	if err != nil {
		http.Error(w, "Publicacion no encontrado", http.StatusBadRequest)
		return
	}
	OpenFile, err := os.Open("uploads/publicaciones/" + pub.Foto)
	if err != nil {
		http.Error(w, "Foto no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar foto", http.StatusBadRequest)
	}
}
