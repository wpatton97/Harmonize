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
