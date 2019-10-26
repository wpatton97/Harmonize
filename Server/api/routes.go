package routes

import (
	"api/indexroute"
	"fmt"
	"net/http"
)

func b() {
	fmt.Println("hello")
}

func Routes() map[string]*http.Handler {
	allRoutes := make(map[string]*http.Handler)

	allRoutes["/"] = indexroute.Index()

	return allRoutes
}
