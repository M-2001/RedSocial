package routers

import (
	"net/http"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*DeleteRelacion realiza el borrado de la relacion de los usuarios*/
func DeleteRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.RelationCollection
	t.UserID = IDUsuario
	t.UserRelacionID = ID

	status, err := bd.DeleteRelacion(t)
	if err != nil || status == false {
		http.Error(w, "Ocurrio un error al intentae elimnar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
