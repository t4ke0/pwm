package sqlite

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

const DB string = "./server/database.db"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", DB)
	checkError(err)
	if db == nil {
		panic("db is nil")
	}
	return db
}

func CreateTables(db *sql.DB) (bool, error) {

	sqlStmt := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255),
			password VARCHAR(100),
			email VARCHAR(255)
		);

	
		CREATE TABLE IF NOT EXISTS passwords(
			pwid INTEGER PRIMARY KEY AUTOINCREMENT,
			user VARCHAR(255),
			passw VARCHAR(100),
			category VARCHAR(25),
			userid INTEGER,
			FOREIGN KEY(userid) REFERENCES users(id)
		);
	`
	_, err := db.Exec(sqlStmt)
	if err == nil {
		return false, nil
	}

	return true, err

}

func CheckForUser(user string, db *sql.DB) int {

	rows, err := db.Query("SELECT username FROM users")
	checkError(err)

	var username string
	var users []string

	for rows.Next() {
		err = rows.Scan(&username)
		checkError(err)
		if user == username {
			users = append(users, username)
		}
	}
	return len(users)
}

func Register(user string, passw string, email string, db *sql.DB) int {

	var rslt int

	if isexist := CheckForUser(user, db); isexist == 1 {
		rslt = 1
		db.Close()
	} else {
		//Hash the password then save it into the db
		h := sha256.New()
		h.Write([]byte(passw))
		ph := h.Sum(nil)
		phx := strings.ToLower(base32.HexEncoding.EncodeToString(ph))
		stmt, err := db.Prepare("INSERT INTO users(username,password,email) VALUES (?,?,?)")
		stmt.Exec(user, phx, email)
		checkError(err)
		rslt = 0
		db.Close()
	}
	return rslt
}

func Login(user string, passw string, db *sql.DB) bool {

	rows, err := db.Query("SELECT * FROM users")
	checkError(err)

	var username string
	var password string
	var email string
	var id int
	var result bool

	h := sha256.New()
	h.Write([]byte(passw))
	ph := h.Sum(nil)
	phx := strings.ToLower(base32.HexEncoding.EncodeToString(ph))

	for rows.Next() {
		err = rows.Scan(&id, &username, &password, &email)
		checkError(err)
		if username == user && password == phx {
			result = true
		} else {
			continue
			//result = false
		}
	}

	rows.Close()
	return result
}

func GetStuff(uid int, category string, db *sql.DB) ([]string, []string, []string, []string) {

	rows, err := db.Query("SELECT pwid ,user,passw,category,userid FROM passwords")
	checkError(err)
	var (
		userid int
		pwid   string
		user   string
		passw  string
		catg   string
		i      []string
		u      []string
		p      []string
		c      []string
	)

	if category == "" {
		for rows.Next() {
			err = rows.Scan(&pwid, &user, &passw, &catg, &userid)
			checkError(err)
			if uid == userid {
				i = append(i, pwid)
				u = append(u, user)
				p = append(p, passw)
				c = append(c, catg)
			}
		}
		rows.Close()
		//		return i, u, p, c
	} else if category != "" {
		for rows.Next() {
			err = rows.Scan(&pwid, &user, &passw, &catg, &userid)
			checkError(err)
			if uid == userid && category == catg {
				i = append(i, pwid)
				u = append(u, user)
				p = append(p, passw)
				c = append(c, catg)
			}
		}
		rows.Close()
		//		return i, u, p, c
	}
	return i, u, p, c
}

func GetUid(user string, db *sql.DB) int {

	rows, err := db.Query("SELECT * FROM users")
	checkError(err)

	var uid int
	var username string
	var password string
	var email string
	var id int
	for rows.Next() {
		err = rows.Scan(&uid, &username, &password, &email)
		checkError(err)
		if username == user {
			id = uid
		}
	}
	rows.Close()
	return id
}

func Update(id int, db *sql.DB, args ...string) []int {

	defer db.Close()
	var f int
	var f0 int
	var f1 int
	var ff []int

	for i, _ := range args {
		if id != 0 && args[i] != "" && i == 0 {
			stmt, err := db.Prepare("UPDATE passwords SET user=? WHERE pwid = ?")
			checkError(err)
			stmt.Exec(args[i], id)
			f = 1
			ff = append(ff, f)

		} else if id != 0 && args[i] != "" && i == 1 {
			stmt, err := db.Prepare("UPDATE passwords SET passw=? WHERE pwid = ?")
			checkError(err)
			stmt.Exec(args[i], id)
			f0 = 1
			ff = append(ff, f0)

		} else if id != 0 && args[i] != "" && i == 2 {
			stmt, err := db.Prepare("UPDATE passwords SET category=? WHERE pwid = ?")
			checkError(err)
			stmt.Exec(args[i], id)
			f1 = 1
			ff = append(ff, f1)

		} else {
			f = 0
			ff = append(ff, f)

		}
	}
	return ff
}

func Save(user string, passwd string, category string, uid int, db *sql.DB) bool {

	var ok bool

	stmt, err := db.Prepare("INSERT INTO passwords(user,passw,category,userid) VALUES (?,?,?,?)")
	stmt.Exec(user, passwd, category, uid)

	if err != nil {
		ok = false
	} else {
		ok = true
	}
	db.Close()
	return ok
}

func Delete(id int, db *sql.DB) bool {
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM passwords WHERE pwid = ?")
	stmt.Exec(id)
	var ok bool
	if err != nil {
		ok = false
	} else {
		ok = true
	}
	return ok
}
