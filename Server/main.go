package main

import (
	"fmt"
	"hackathon/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = 8080

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	GETroutes := routes.GETRoutes()
	for path, handler := range GETroutes {
		api.HandleFunc(path, handler).Methods("GET")
	}

	POSTroutes := routes.POSTRoutes()
	for path, handler := range POSTroutes {
		api.HandleFunc(path, handler).Methods("POST")
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Printf("Listening on port: %d", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS()(router))

}
