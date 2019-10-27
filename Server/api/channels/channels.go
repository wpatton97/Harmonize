package channels

import (
	"encoding/json"
	"hackathon/db"
	"net/http"
)

const staticPath string = "./static/"

func Get(w http.ResponseWriter, r *http.Request) {

	channels := db.GetChannel()
	data, _ := json.Marshal(&channels)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
