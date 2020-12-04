package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/M-2001/RedSocial/bd"
)

/*ReadPublicacionesSeguidores permitira leer las publicaciones de los seguidores*/
func ReadPublicacionesSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina mayor a cero", http.StatusBadRequest)
		return
	}
	resp, correct := bd.ReadPublicacionesSeguidores(IDUsuario, page)
	// for i := range resp {
	// 	OpenFile, err := os.Open("uploads/publicaciones/" + resp[i].Publicacion.Foto)
	if correct == false {
		http.Error(w, "Error a leer el contenido", http.StatusBadRequest)
		return
	}
	// _, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar foto", http.StatusBadRequest)
	}
	// }
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
