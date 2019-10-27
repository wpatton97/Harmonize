package candidates

import (
	"fmt"
	"hackathon/db"
	"net/http"
)

const staticPath string = "./static/"

func Get(w http.ResponseWriter, r *http.Request) {
	template := `
		{
			"Song":{
				"ID": '%d',
				"Title": "%s",
				"Author": "%s",
				"Album": "%s",
				"Art": "%s",
				"Length": '%d'
			},
			"Addedby":{
				"Name": "%s",
				"Avatar": "%s"
			},
			"Votes": ["%s","%s"]
		}
	`
	js := "["
	for i := 0; i <= 6; i++ {
		curSong := db.GetRandomSong()
		js += fmt.Sprintf(template, curSong.ID, curSong.Title, curSong.Author, curSong.Album, curSong.Art, curSong.Length, "xinx", "/userprofiles/xinx.png", "/userprofiles/xinx.png", "/userprofiles/goose.png")
		if i != 6 {
			js += ","
		} else {
			js += "]"
		}
	}
	/*	js := `
				[
				{
					"Song":{
					   	"ID": 1,
						"Title": "Left Hand Free",
						"Author": "alt-J",
						"Album": "This Is All Yours",
						"Art": "https://upload.wikimedia.org/wikipedia/en/0/0c/Alt-J_-_This_is_all_yours.jpg",
						"Length": 120
					},
					"Addedby": {
				    	"Name": "goose",
				    	"Avatar": "https://images-ext-1.discordapp.net/external/0M2bB5sJiQyp57Es--yF6ZM5IQgW0zts9fblg-SRhzc/https/i.imgur.com/Ska1zn6.png?width=247&height=301"
					},
					"Votes": ["https://place-hold.it/32" , "https://place-hold.it/32" ]
				},
				{
					"Song":{
					   	"ID": 2,
						"Title": "Deadcrush",
						"Author": "alt-J",
						"Album": "RELAXER",
						"Art": "https://upload.wikimedia.org/wikipedia/en/d/dc/Alt-J_-_Relaxer_cover.jpg" ,
						"Length": 42069
					    },
					    "Addedby": {
					    	"Name": "xinx",
					    	"Avatar": "https://place-hold.it/32"
					    },
					"Votes": ["https://place-hold.it/32" , "https://place-hold.it/32" ]
				},
				{
					"Song":{
					   	"ID": 3,
						"Title": "Adeline",
						"Author": "alt-J",
						"Album": "RELAXER",
						"Art": "https://upload.wikimedia.org/wikipedia/en/d/dc/Alt-J_-_Relaxer_cover.jpg" ,
						"Length": 42069
					    },
					    "Addedby": {
					    	"Name": "xinx",
					    	"Avatar": "https://place-hold.it/32"
					    },
					"Votes": ["https://place-hold.it/32" , "https://place-hold.it/32" ]
				}
			]
				`

				js2 := `
				[
				{
					"Song":{
					   	"ID": 2,
						"Title": "Deadcrush",
						"Author": "alt-J",
						"Album": "RELAXER",
						"Art": "https://upload.wikimedia.org/wikipedia/en/d/dc/Alt-J_-_Relaxer_cover.jpg" ,
						"Length": 42069
					    },
					    "Addedby": {
					    	"Name": "xinx",
					    	"Avatar": "https://place-hold.it/32"
					    },
					"Votes": ["https://place-hold.it/32" , "https://place-hold.it/32" ]
				},
				{
					"Song":{
					   	"ID": 3,
						"Title": "Adeline",
						"Author": "alt-J",
						"Album": "RELAXER",
						"Art": "https://upload.wikimedia.org/wikipedia/en/d/dc/Alt-J_-_Relaxer_cover.jpg" ,
						"Length": 42069
					    },
					    "Addedby": {
					    	"Name": "xinx",
					    	"Avatar": "https://place-hold.it/32"
					    },
					"Votes": ["https://place-hold.it/32" , "https://place-hold.it/32" ]
				}
			]
			`

		keys, ok := r.URL.Query()["channel"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'channel' is missing")
			return
		}
		channelid, _ := strconv.Atoi(keys[0])
	*/
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(js))

}
