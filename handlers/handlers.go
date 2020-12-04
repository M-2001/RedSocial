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

	//Insert Publication Json
	router.HandleFunc("/publicacionJ", middlew.ChequeoBD(middlew.ValidacionJWT(routers.GrabarPublicacionJSON))).Methods("POST")

	router.HandleFunc("/readPublicacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadPublicaciones))).Methods("GET")
	router.HandleFunc("/deletePublicacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.DeletePublicacion))).Methods("DELETE")
	/*EndPoints para imagenes*/
	router.HandleFunc("/upAvatar", middlew.ChequeoBD(middlew.ValidacionJWT(routers.UpAvatar))).Methods("POST")
	router.HandleFunc("/upBanner", middlew.ChequeoBD(middlew.ValidacionJWT(routers.UpBanner))).Methods("POST")
	router.HandleFunc("/upImagePublicacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.UpImagePublicacion))).Methods("POST")
	router.HandleFunc("/mostrarAvatr", middlew.ChequeoBD(routers.MostrarAvatar)).Methods("GET")
	router.HandleFunc("/mostrarBanner", middlew.ChequeoBD(routers.MostrarBanner)).Methods("GET")
	/*EndPoint para insertar una relacion*/
	router.HandleFunc("/insertRelation", middlew.ChequeoBD(middlew.ValidacionJWT(routers.Relation))).Methods("POST")
	/*EndPoint eliminar relacion*/
	router.HandleFunc("/unfollow", middlew.ChequeoBD(middlew.ValidacionJWT(routers.DeleteRelacion))).Methods("DELETE")
	/*mostrar relacion*/
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ConsultaRelacion))).Methods("GET")
	/*mostrar usuarios relacionados*/
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ListUsers))).Methods("GET")
	/*todas las publicaciones de los seguidores*/
	router.HandleFunc("/allPublicaciones", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadPublicacionesSeguidores))).Methods("GET")
	//Mostrar Una publicacion
	router.HandleFunc("/Apublication", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadUnaPublicacion))).Methods("GET")

	//mostrar Foto Publicacion
	router.HandleFunc("/mostrarFotoPub", middlew.ChequeoBD(routers.MostrarFotoPublicacion)).Methods("GET")

	// Rutas Comentarios
	router.HandleFunc("/comentar", middlew.ChequeoBD(middlew.ValidacionJWT(routers.HacerComentario))).Methods("POST")
	//mostrar comentarios
	router.HandleFunc("/mostrarComentarios", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadCometarios))).Methods("GET")
	//eliminar comentario
	router.HandleFunc("/deleteComentario", middlew.ChequeoBD(middlew.ValidacionJWT(routers.DeleteComentario))).Methods("DELETE")

	// Rutas reaccion
	router.HandleFunc("/reaccion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReaccionPublicacion))).Methods("POST")
	router.HandleFunc("/mostrarReacciones", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadReacciones))).Methods("GET")
	router.HandleFunc("/delreaccion", middlew.ChequeoBD(middlew.ValidacionJWT(routers.DeleteReaccion))).Methods("DELETE")
	//mostrar todas las reacciones
	router.HandleFunc("/AllReacciones", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadReaccionesDeUnaPublicacion))).Methods("GET")

	//Reacciones en los comentarios
	router.HandleFunc("/reactComment", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReactComment))).Methods("POST")

	router.HandleFunc("/readRComment", middlew.ChequeoBD(middlew.ValidacionJWT(routers.ReadReactComment))).Methods("GET")

	router.HandleFunc("/delRComment", middlew.ChequeoBD(middlew.ValidacionJWT(routers.DeleteReactComment))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
