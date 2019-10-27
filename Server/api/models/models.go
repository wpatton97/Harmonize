package models

import (
	"net/http"
	"time"
)

const staticPath string = "./static/"

type UserModel struct {
	ID             int
	Name           string
	ProfilePicture string
	Cookie         string
	ActiveChannel  int
}

type ChannelModel struct {
	ID   int
	Name string
	//VoteID int
}

type SongModel struct {
	Title    string
	Author   string
	Album    string
	Art      string
	Length   int
	ID       int
	Location string
}

type UserVotes struct {
	UserID   int
	VoteID   int
	DateTime time.Time
}

type Votes struct {
	ID            int
	SongID        int
	ChannelID     int
	Votes         int
	Completed     int
	DateTime      time.Time
	InitiatedUser int
}

func Get(w http.ResponseWriter, r *http.Request) {

	js := "{\"test\":\"Werks\"}"
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}
