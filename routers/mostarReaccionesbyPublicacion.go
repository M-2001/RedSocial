package routers

import (
	"encoding/json"
	"net/http"

	"github.com/M-2001/RedSocial/bd"
)

/*ReadReacciones servira para leer todas las de una publicacion*/
func ReadReacciones(w http.ResponseWriter, r *http.Request) {
	IDP := r.URL.Query().Get("id")
	if len(IDP) < 1 {
		http.Error(w, "Debe enviar el parametro id para ver reacciones de la publicacion", http.StatusBadRequest)
		return
	}
	/*if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviarse el parametro pagina", http.StatusBadRequest)
		return
	}*/
	/*page, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviarse el parametro pagina   con un valor mayor a cero", http.StatusBadRequest)
		return
	}*/
	//pag := int64(page)

	respuesta, correct := bd.ReadReacciones(IDP)
	if correct == false {
		http.Error(w, "Error a leer el contenido!!! Intente nuevamente", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
