package routers

import (
	"encoding/json"
	"net/http"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/models"
)

/*Registro es la funcion que permite agregar usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) > 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)

	}
	_, encontrado, _ := bd.ChequeoYaExiste(t.Email)
	if encontrado == true {
		http.Error(w, "ya existe un registro con este correo", 400)
		return
	}
	_, status, err := bd.InsertRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error, intente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro guardar el registro", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
