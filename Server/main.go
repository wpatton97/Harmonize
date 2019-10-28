package main

import (
	//"fmt"
	"hackathon/routes"
	"log"
	"net/http"
	"golang.org/x/crypto/acme/autocert"
	"crypto/tls"
	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = 80

func main() {

	certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("ragex04.dev"), //Your domain here
        Cache:      autocert.DirCache("/etc/letsencrypt/live/ragex04.dev"),            //Folder for storing certificates
    }

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

	server := &http.Server{
        Addr: ":https",
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
		},
		Handler: router,
    }

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	log.Printf("Listening on port: %d", port)

	log.Fatal(server.ListenAndServeTLS("", "")) //Key and cert are coming from Let's Encrypt
	// err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "/etc/letsencrypt/live/ragex04.dev/fullchain.pem", "/etc/letsencrypt/live/ragex04.dev/privkey.pem", handlers.CORS()(router))
	// if err != nil{
	// 	log.Fatal(err)
	// }

}
