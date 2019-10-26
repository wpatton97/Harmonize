module indexroute

import (
	"net/http"
)

const staticPath string = "./static/"

func (h http.HandlerFunc) Index(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir(staticPath))
	fs.ServeHTTP(w, r)
}
