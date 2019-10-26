package routes

import (
	"api/indexroute"
	"net/http"
)

func Routes() map[string]*http.Handler {
	allRoutes := make(map[string]*http.Handler)

	allRoutes["/"] = indexroute.Index()

	return allRoutes
}
