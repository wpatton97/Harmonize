package routes

import (
	"hackathon/api/available"
	"hackathon/api/channels"
	"hackathon/api/chat"
	"hackathon/api/nowplaying"
	"hackathon/api/users"
	"hackathon/api/vote"
	"net/http"
)

func GETRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	allRoutes := make(map[string]func(http.ResponseWriter, *http.Request))

	allRoutes["/available"] = available.Get
	allRoutes["/channels"] = channels.Get
	allRoutes["/nowplaying"] = nowplaying.Get
	allRoutes["/users"] = users.Get
	allRoutes["/vote"] = vote.Get
	allRoutes["/chat"] = chat.Get
	return allRoutes
}

func POSTRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	allRoutes := make(map[string]func(http.ResponseWriter, *http.Request))

	allRoutes["/vote"] = vote.Post

	return allRoutes
}
