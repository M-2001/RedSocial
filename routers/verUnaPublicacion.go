package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/M-2001/RedSocial/bd"
)

/*ReadUnaPublicacion servira para leer todas las publicaciones*/
func ReadUnaPublicacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("idPublicacion")
	if len(ID) < 1 {
		http.Error(w, "Debe enviarse el parametro id", http.StatusBadRequest)
		return
	}
	/*if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviarse el parametro pagina", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviarse el parametro pagina   con un valor mayor a cero", http.StatusBadRequest)
		return
	}*/
	//pag := int64(page)
	respuesta, correct := bd.ReadUnaPublicacion(ID)
	for i := range respuesta {
		OpenFile, err := os.Open("uploads/publicaciones/" + respuesta[i].Foto)

		if correct == false {
			http.Error(w, "Error a leer el contenido", http.StatusBadRequest)
			return
		}
		_, err = io.Copy(w, OpenFile)
		if err != nil {
			http.Error(w, "Error al copiar foto", http.StatusBadRequest)
		}

	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
