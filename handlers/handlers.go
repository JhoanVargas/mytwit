package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JhoanVargas/mytwit/middlew"
	"github.com/JhoanVargas/mytwit/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/twit", middlew.ChequeoBD(middlew.ValidoJWT(routers.GrabarTwit))).Methods("POST")
	router.HandleFunc("/leerTwits", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeerTwits))).Methods("GET")
	router.HandleFunc("/eliminarTwit", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTwit))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerrAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//Permisos
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
