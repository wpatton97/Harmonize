package nowplaying

import (
	"net/http"
	"hackathon/db"
)

const staticPath string = "./static/"

func Get(w http.ResponseWriter, r *http.Request) {

	js := "{\"test\":\"Werks\"}"
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
