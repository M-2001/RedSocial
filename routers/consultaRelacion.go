package routers

import (
	"encoding/json"
	"net/http"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*ConsultaRelacion se encarga de ver si esxiste relacion*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.RelationCollection
	t.UserID = IDUsuario
	t.UserRelacionID = ID

	var resp models.ConsultaRelacion
	status, err := bd.ConsultaRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
