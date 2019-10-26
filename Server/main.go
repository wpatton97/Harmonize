package main

import (
	"fmt"
	"hackathon/routes"
	"log"
<<<<<<< HEAD

	. "github.com/Ragex04/Harmonize/Server/api/indexroute.go"
	. "github.com/Ragex04/Harmonize/Server/api/routes.go"
=======
	"net/http"

	"github.com/gorilla/mux"
>>>>>>> master
)

const port = 8080

func main() {
<<<<<<< HEAD
	log.Printf("YEET\n")
	log.Printf(routes.b())
=======
	router := mux.NewRouter()
	routes := routes.Routes()
	for path, handler := range routes {
		router.HandleFunc(path, handler)
	}

	log.Printf("Listening on port: %d", 8080)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)

>>>>>>> master
}
