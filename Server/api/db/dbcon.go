package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase()(*sql.DB, error){
	serverName := "35.202.159.22"
	user := "root"
	password := "o,|D-Q#[;&=Xznq#>C6Wa=w7*jM[iq:\\R'K|r-qatWXW1WJ{"
	dbName := "Harmonizer"

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", conn)
	if err != nil{
		return nil, err
	}
	return db, nil
}