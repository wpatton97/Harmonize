package main

import (
	"fmt"
	"hackathon/api/db"
	"hackathon/routes"
	"log"
	"net/http"
)

const port = 8080

func postSong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("attempting to insert into db")
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("db conn failed")
	}
	_, err = database.Exec("INSERT INTO 'songs' VALUES ('/docs/music/intro.mp3', 'alt-j', 'an awesome wave', '/static/alt-j.jpg')")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	//router := mux.NewRouter()
	routes := routes.Routes()
	for path, handler := range routes {
		router.HandleFunc(path, handler)
	}

	log.Printf("Listening on port: %d", 8080)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)

}
