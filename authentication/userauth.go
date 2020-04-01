package authentication

import (
	"database/sql"

	"../sqlite"
)

// DB const variable indicates the path of the sqlite3 file
const DB string = "./server/database.db"

// Register Saves New Users into Sqlite DB
func Register(username string, password string, email string) bool {
	db, err := sql.Open("sqlite3", DB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var ok bool
	reg := sqlite.Register(username, password, email, db)
	if reg == 0 {
		ok = true
	} else {
		ok = false
	}
	return ok
}

// Login Checks if the username and password entered are the same as others in the DB
func Login(username string, password string) bool {
	db, err := sql.Open("sqlite3", DB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var ok bool
	reg := sqlite.Login(username, password, db)
	if reg {
		ok = true
	} else {
		ok = false
	}
	return ok
}
