package routers

import (
	"encoding/json"
	"net/http"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*ReactComment permitira realizar registro de reacciones a comentarios */
func ReactComment(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("commentid")

	reactComment := r.URL.Query().Get("react")

	if len(ID) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	if reactComment > "2" {
		http.Error(w, "reactComment debe ser menor igual a 2 ", http.StatusBadRequest)
		return
	} else if reactComment < "1" {
		http.Error(w, "reactComment debe ser igual o mayor a 1 ", http.StatusBadRequest)
		return
	}

	var rctn models.ReactComment
	err := json.NewDecoder(r.Body).Decode(&rctn)
	reaccion := models.ReactComment{
		UserID:       IDUsuario,
		CommentID:    ID,
		ReactComment: reactComment,
	}

	_, status, err := bd.ReactComment(reaccion)
	if err != nil {
		http.Error(w, "Ocurrio un error inesperado! intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Ocurrio un error inesperado! no se pudo insertar reaccion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reaccion)
}

/*ReadReactComment servira para leer todas las de una publicacion*/
func ReadReactComment(w http.ResponseWriter, r *http.Request) {
	IDp := r.URL.Query().Get("idP")
	if len(IDp) < 1 {
		http.Error(w, "Debe enviar el parametro id para ver reacciones de los comentarios", http.StatusBadRequest)
		return
	}

	respuesta, correct := bd.ReadReactComment(IDp)
	if correct == false {
		http.Error(w, "Error a leer el contenido!!! Intente nuevamente", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

/*DeleteReactComment servira para eliminar una publicacion*/
func DeleteReactComment(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviarse el parametro id", http.StatusBadRequest)
		return
	}
	err := bd.DeleteReactComment(ID, IDUsuario)
	if err != nil {
		http.Error(w, "No se pudo eliminar la reaccion Intente mas tarde"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
