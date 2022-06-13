package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/luizclaudioholanda/go-rest-api/controllers"
	"github.com/luizclaudioholanda/go-rest-api/middleware"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.ListaPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidadePorId).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8888", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
