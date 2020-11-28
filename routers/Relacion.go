package routers

import (
	"net/http"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*Relation permitira realizar registro de relacion*/
func Relation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var relation models.RelationCollection
	relation.UserID = IDUsuario
	relation.UserRelacionID = ID

	status, err := bd.InsertRelation(relation)
	if err != nil {
		http.Error(w, "Ocurrio un error inesperado! intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Ocurrio un error inesperado! no se pudo insertar relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
