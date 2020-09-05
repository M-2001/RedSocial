package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/M-2001/RedSocial/middlew"
	"github.com/M-2001/RedSocial/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto, el handler y pongo en ejecucion el puerto
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.LogIn)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidacionJWT(routers.VerProfile))).Methods("GET")
	router.HandleFunc("/updateperfil", middlew.ChequeoBD(middlew.ValidacionJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/publicacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.GrabarPublicacion))).Methods("POST")
	router.HandleFunc("/readPublicacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadPublicaciones))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
