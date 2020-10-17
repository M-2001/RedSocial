package routers

import (
	"encoding/json"
	"net/http"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*ReaccionPublicacion permitira realizar registro de reaccion*/
func ReaccionPublicacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("publicacionid")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var rctn models.Reaccion
	err := json.NewDecoder(r.Body).Decode(&rctn)
	reaccion := models.ReaccionCollection{
		UserID:        IDUsuario,
		PublicacionID: ID,
	}

	/*var reaccion models.ReaccionCollection
	reaccion.UserID = IDUsuario
	reaccion.PublicacionID = ID*/

	_, status, err := bd.ReaccionPublicacion(reaccion)
	if err != nil {
		http.Error(w, "Ocurrio un error inesperado! intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Ocurrio un error inesperado! no se pudi insertar reaccion"+err.Error(), http.StatusBadRequest)
		return
	}
	/*w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)*/
	json.NewEncoder(w).Encode(reaccion)
}
