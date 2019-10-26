package routes

import (
	"hackathon/api/index"
	"net/http"
)

func Routes() map[string]func(http.ResponseWriter, *http.Request) {
	allRoutes := make(map[string]func(http.ResponseWriter, *http.Request))

	allRoutes["/"] = index.Index

	return allRoutes
}
