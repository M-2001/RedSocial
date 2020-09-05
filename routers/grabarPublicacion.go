package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*GrabarPublicacion sirve para grabar un publicacion en la base de datos*/
func GrabarPublicacion(w http.ResponseWriter, r *http.Request) {
	var msj models.Publicacion
	err := json.NewDecoder(r.Body).Decode(&msj)

	registro := models.GraboPublicacion{
		UserId:           IDUsuario,
		Publicacion:      msj.Publicacion,
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
	w.WriteHeader(http.StatusCreated)
}
