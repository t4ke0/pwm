package authentication

import (
	"../sqlite"
	"database/sql"
)

const DB string = "./server/database.db"

func Register(username string, password string, email string) bool {
	db, err := sql.Open("sqlite3", DB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var ok bool
	//db := sqlite.InitDb()
	reg := sqlite.Register(username, password, email, db)
	if reg == 0 {
		ok = true
	} else {
		ok = false
	}
	return ok
}

func Login(username string, password string) bool {
	db, err := sql.Open("sqlite3", DB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var ok bool
	//db := sqlite.InitDb()
	reg := sqlite.Login(username, password, db)
	if reg {
		ok = true
	} else {
		ok = false
	}
	return ok
}
