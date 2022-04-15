package forum

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type id struct {
	email    string
	username string
	password string
}

func StartDB() {
	db, err := sql.Open("sqlite3", "data_base.db")
	if err != nil {
		log.Fatal(err)
	}
	Table := `
	CREATE TABLE type (
		id	TEXT
	)
	`
	db.Exec(Table)
}

func SendToDBRegister(email string, username string, password string) {
	db, err := sql.Open("sqlite3", "data_base.db")
	if err != nil {
		log.Fatal(err)
	}
	email = RemoveFromString(email)
	username = RemoveFromString(username)
	password = RemoveFromString(password)
	db.Exec("INSERT INTO profile-data VALUES(?,?,?);", email, username, password)
}

func RemoveFromString(arg string) string {
	arg = strings.ReplaceAll(arg, "[", "")
	arg = strings.ReplaceAll(arg, "]", "")
	return arg
}
