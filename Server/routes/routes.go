package routes

import (
	"hackathon/api/available"
	"hackathon/api/channels"
	"hackathon/api/chat"
	"hackathon/api/index"
	"hackathon/api/nowplaying"
	"hackathon/api/users"
	"hackathon/api/vote"
	"net/http"
)

func Routes() map[string]func(http.ResponseWriter, *http.Request) {
	allRoutes := make(map[string]func(http.ResponseWriter, *http.Request))

	allRoutes["/"] = index.Index
	allRoutes["/api/available"] = available.Get
	allRoutes["/api/channels"] = channels.Get
	allRoutes["/api/nowplaying"] = nowplaying.Get
	allRoutes["/api/users"] = users.Get
	allRoutes["/api/vote"] = vote.Get
	allRoutes["/api/chat"] = chat.Get
	return allRoutes
}
