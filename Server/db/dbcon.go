package db

import (
	"database/sql"
	"fmt"
	"hackathon/api/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// CreateDatabase sets up a database connection
func CreateDatabase() (*sql.DB, error) {
	serverName := "35.202.159.22"
	user := "root"
	password := "o,|D-Q#[;&=Xznq#>C6Wa=w7*jM[iq:\\R'K|r-qatWXW1WJ{"
	dbName := "Harmonizer"

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func checkerr(err error) bool {
	if err != nil {
		log.Panicln(err.Error())
		return false
	} else {
		return true
	}
}

// GetSongs Grabs all songs from database
func GetSongs() []models.SongModel {
	var songs []models.SongModel

	db, err := CreateDatabase()
	checkerr(err)

	rows, err := db.Query("SELECT * FROM Songs")
	checkerr(err)

	for rows.Next() {
		song := models.SongModel{}
		err = rows.Scan(&song.ID, &song.Title, &song.Author, &song.Album, &song.Art, &song.Length, &song.Location)
		if !checkerr(err) {
			continue
		}

		songs = append(songs, song)
	}

	db.Close()
	return songs
}

// GetRandomSong Exactly what it says
func GetSongByID(id int) models.SongModel {
	var songs []models.SongModel

	db, err := CreateDatabase()
	checkerr(err)

	rows, err := db.Query("SELECT * FROM Songs WHERE ID=%d", id)
	checkerr(err)

	for rows.Next() {
		song := models.SongModel{}
		err = rows.Scan(&song.ID, &song.Title, &song.Author, &song.Album, &song.Art, &song.Length, &song.Location)
		if !checkerr(err) {
			continue
		}

		songs = append(songs, song)
	}

	db.Close()
	return songs[0]
}

// GetRandomSong Exactly what it says
func GetRandomSong() models.SongModel {
	var songs []models.SongModel

	db, err := CreateDatabase()
	checkerr(err)

	rows, err := db.Query("SELECT * FROM Songs ORDER BY RAND() LIMIT 1")
	checkerr(err)

	for rows.Next() {
		song := models.SongModel{}
		err = rows.Scan(&song.ID, &song.Title, &song.Author, &song.Album, &song.Art, &song.Length, &song.Location)
		if !checkerr(err) {
			continue
		}

		songs = append(songs, song)
	}

	db.Close()
	return songs[0]
}

// GetSongsLike Grabs all songs from database LIKE %query%
func GetSongsLike(searchterm string) []models.SongModel {
	var songs []models.SongModel

	db, err := CreateDatabase()
	checkerr(err)

	query := `
		SELECT * FROM Songs WHERE Title LIKE '%%%s%%' OR Author LIKE '%%%s%%';
	`
	query = fmt.Sprintf(query, searchterm)
	rows, err := db.Query(query)
	checkerr(err)

	for rows.Next() {
		song := models.SongModel{}
		err = rows.Scan(&song.ID, &song.Title, &song.Author, &song.Album, &song.Art, &song.Length, &song.Location)
		if !checkerr(err) {
			continue
		}

		songs = append(songs, song)
	}

	db.Close()
	return songs
}

// gets all users in db
func GetUsers() []models.UserModel {
	var users []models.UserModel

	db, err := CreateDatabase()
	checkerr(err)

	rows, err := db.Query("SELECT * FROM Users")
	checkerr(err)

	for rows.Next() {
		user := models.UserModel{}
		err = rows.Scan(&user.ID, &user.Name, &user.ProfilePicture, &user.Cookie, &user.ActiveChannel)
		if !checkerr(err) {
			continue
		}
		users = append(users, user)
	}
	db.Close()
	return users
}

//gets all users containing searchterm
func GetUsersLike(searchterm string) []models.UserModel {
	var users []models.UserModel
	db, err := CreateDatabase()
	checkerr(err)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM Users WHERE Name like '%%%s%%'", searchterm))
	checkerr(err)

	for rows.Next() {
		user := models.UserModel{}
		err = rows.Scan(&user.ID, &user.Name, &user.ProfilePicture, &user.Cookie, &user.ActiveChannel)
		if !checkerr(err) {
			continue
		}
		users = append(users, user)
	}
	db.Close()
	return users
}

//gets all channels
func GetChannel() []models.ChannelModel {
	var channels []models.ChannelModel
	db, err := CreateDatabase()
	checkerr(err)
	rows, err := db.Query("SELECT * FROM Channels")
	checkerr(err)
	for rows.Next() {
		channel := models.ChannelModel{}
		err = rows.Scan(&channel.ID, &channel.Name)
		if !checkerr(err) {
			continue
		}
		channels = append(channels, channel)
	}
	db.Close()
	return channels
}

//gets all channels that contain searchterm
func GetChannelLike(searchterm string) []models.ChannelModel {
	var channels []models.ChannelModel
	db, err := CreateDatabase()
	checkerr(err)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM Channels WHERE Name LIKE '%%%s%%' ", searchterm))
	checkerr(err)
	for rows.Next() {
		channel := models.ChannelModel{}
		err = rows.Scan(&channel.ID, &channel.Name)
		if !checkerr(err) {
			continue
		}
		channels = append(channels, channel)
	}
	db.Close()
	return channels
}

//get all users votes
func GetUserVotes() []models.UserVotes {
	var uservotes []models.UserVotes
	db, err := CreateDatabase()
	checkerr(err)
	rows, err := db.Query("SELECT * FROM UserVotes")
	checkerr(err)
	for rows.Next() {
		uservote := models.UserVotes{}
		err = rows.Scan(&uservote.UserID, &uservote.VoteID, &uservote.DateTime)
		if !checkerr(err) {
			continue
		}
		uservotes = append(uservotes, uservote)
	}
	db.Close()
	return uservotes
}

//gets all users votes contained in searchterm
func GetUserVotesLikes(searchterm int) []models.UserVotes {
	var uservotes []models.UserVotes
	db, err := CreateDatabase()
	checkerr(err)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM UserVotes WHERE UserID LIKE '%%%d%%'", searchterm))
	checkerr(err)
	for rows.Next() {
		uservote := models.UserVotes{}
		err = rows.Scan(&uservote.UserID, &uservote.VoteID, &uservote.DateTime)
		if !checkerr(err) {
			continue
		}
		uservotes = append(uservotes, uservote)
	}
	db.Close()
	return uservotes
}

//gets all Votes
func GetVotes() []models.Votes {
	var votes []models.Votes
	db, err := CreateDatabase()
	checkerr(err)
	rows, err := db.Query("SELECT * FROM Votes")
	checkerr(err)
	for rows.Next() {
		vote := models.Votes{}
		err = rows.Scan(&vote.ID, &vote.SongID, &vote.ChannelID, &vote.Votes, &vote.Completed, &vote.DateTime, &vote.InitiatedUser)
		if !checkerr(err) {
			continue
		}
		votes = append(votes, vote)
	}
	db.Close()
	return votes
}

//get all votes contained in searchterm
func GetVotesLike(searchterm int) []models.Votes {
	var votes []models.Votes
	db, err := CreateDatabase()
	checkerr(err)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM Votes WHERE ID LIKE '%%%d%%'", searchterm))
	checkerr(err)
	for rows.Next() {
		vote := models.Votes{}
		err = rows.Scan(&vote.ID, &vote.SongID, &vote.ChannelID, &vote.Votes, &vote.Completed, &vote.DateTime, &vote.InitiatedUser)
		if !checkerr(err) {
			continue
		}
		votes = append(votes, vote)
	}
	db.Close()
	return votes
}

func GetNowPlaying(channelid int) models.NowPlaying {
	var Np models.NowPlaying

	var votes []models.Votes

	db, err := CreateDatabase()
	checkerr(err)

	q := "SELECT * FROM Votes WHERE ChannelID=%d AND Completed=1 ORDER BY DateTime"
	q = fmt.Sprintf(q, channelid)

	rows, err := db.Query(q)
	checkerr(err)

	for rows.Next() {
		var v models.Votes
		rows.Scan(&v.ID, &v.SongID, &v.ChannelID, &v.Votes, &v.Completed, &v.DateTime, &v.InitiatedUser)
		votes = append(votes, v)
	}

	if len(votes) == 0 {
		song := GetRandomSong()
		Np.Title = song.Title
		Np.Author = song.Author
		Np.Art = song.Art
		Np.Album = song.Album
		Np.URL = song.Location
		Np.Time.Length = song.Length
		Np.Time.Current = 0
	} else {
		song := GetSongByID(votes[0].SongID)
		Np.Title = song.Title
		Np.Author = song.Author
		Np.Art = song.Art
		Np.Album = song.Album
		Np.URL = song.Location
		Np.Time.Length = song.Length
		Np.Time.Current = 0
	}

	return Np
}
