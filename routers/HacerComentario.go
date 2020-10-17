package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*HacerComentario sirve para agregar un coemtario a publicacion en la base de datos*/
func HacerComentario(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("publicacionid")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var cmtr models.Comentario
	err := json.NewDecoder(r.Body).Decode(&cmtr)
	registro := models.CommentPublications{
		UserID:          IDUsuario,
		PublicacionID:   ID,
		Comentario:      cmtr.Comentario,
		FechaComentario: time.Now(),
	}
	_, status, err := bd.DoCommentPublication(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar hacer comentario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado hacer el comentario", 400)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registro)
}
