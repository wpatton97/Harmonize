package main

import (
	"fmt"
	"hackathon/db"
	"hackathon/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = 8080

func main() {
	router := mux.NewRouter()
	routes := routes.Routes()
	for path, handler := range routes {
		router.HandleFunc(path, handler)
	}
	log.Printf("%v", db.GetSongsLike("Interlude"))
	log.Printf("Listening on port: %d", 8080)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)

}
