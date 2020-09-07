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
	router.HandleFunc("/deletePublicacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.DeletePublicacion))).Methods("DELETE")
	/*EndPoints para imagenes*/
	router.HandleFunc("/upAvatar", middlew.ChequeoBD(middlew.ValidacionJWT(routers.UpAvatar))).Methods("POST")
	router.HandleFunc("/upBanner", middlew.ChequeoBD(middlew.ValidacionJWT(routers.UpBanner))).Methods("POST")
	router.HandleFunc("/mostrarAvatr", middlew.ChequeoBD(routers.MostrarAvatar)).Methods("GET")
	router.HandleFunc("/mostrarBanner", middlew.ChequeoBD(routers.MostrarBanner)).Methods("GET")
	/*EndPoint para insertar una relacion*/
	router.HandleFunc("/insertRelation", middlew.ChequeoBD(middlew.ValidacionJWT(routers.Relation))).Methods("POST")
	/*EndPoint eliminar relacion*/
	router.HandleFunc("/unfollow", middlew.ChequeoBD(middlew.ValidacionJWT(routers.DeleteRelacion))).Methods("DELETE")
	/*mostrar relacion*/
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ConsultaRelacion))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
