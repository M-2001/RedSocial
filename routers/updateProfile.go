package routers

import (
	"encoding/json"
	"net/http"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*UpdateProfile servira para actualizar registros*/
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificarRegistro(user, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar actualizar el registro"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se pudo Actualizar", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
