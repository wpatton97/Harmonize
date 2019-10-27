package nowplaying

import (
	"encoding/json"
	"hackathon/api/models"
	"hackathon/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

const staticPath string = "./static/"

var channel2nowplaying = make(map[int]models.NowPlaying)
var channel2startoffset = make(map[int]time.Time)

func Get(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["channel"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'channel' is missing")
		return
	}
	channelid, _ := strconv.Atoi(keys[0])

	inStor, exist := channel2nowplaying[channelid]
	if exist {
		if inStor.Time.Current > float64(inStor.Time.Length) {
			inStor = db.GetNowPlaying(channelid)
			channel2nowplaying[channelid] = inStor
			channel2startoffset[channelid] = time.Now()
		} else {
			inStor.Time.Current = time.Now().Sub(channel2startoffset[channelid]).Seconds()
		}

	} else {
		log.Println(inStor.Time.Current)
		inStor = db.GetNowPlaying(channelid)
		channel2nowplaying[channelid] = inStor
		channel2startoffset[channelid] = time.Now()
	}

	data, _ := json.Marshal(inStor)
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}
