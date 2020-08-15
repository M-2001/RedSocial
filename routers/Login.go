package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/M-2001/RedSocial/bd"
	"github.com/M-2001/RedSocial/jwt"
	"github.com/M-2001/RedSocial/models"
)

/*LogIn realiza el login Usuario*/
func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario y contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
	}

	doct, present := bd.Login(user.Email, user.Password)
	if present == false {
		http.Error(w, "El email de Usuario se requiere", 400)
		return
	}

	key, err := jwt.GeneroJWT(doct)
	if err != nil {
		http.Error(w, "Error, El token no fue generado"+err.Error(), 400)
		return
	}

	reply := models.ReplyLogin{
		Token: key,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reply)

	cookieTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   key,
		Expires: cookieTime,
	})
}
