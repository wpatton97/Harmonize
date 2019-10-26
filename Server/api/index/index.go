package index

import (
	"net/http"
)

const staticPath string = "./static/"

func Index(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir(staticPath))
	fs.ServeHTTP(w, r)
}
