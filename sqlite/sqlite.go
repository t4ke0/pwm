package sqlite

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"fmt"
	"log"

	// _ blank for importing sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

// DB database file path
const DB string = "./server/database.db"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// InitDb Open sqlite3 File
func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", DB)
	checkError(err)
	if db == nil {
		err := fmt.Errorf("DB is Nil")
		checkError(err)
	}
	return db
}

// CreateTables create tables
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

// CheckForUser check if already exist in the DB & returns the length of users array
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

//CountUsers Count the number of user that we have in the Database
func CountUsers(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT COUNT (*) FROM users;")
	if err != nil {
		return 0, err
	}
	var amount int
	for rows.Next() {
		err = rows.Scan(&amount)
		if err != nil {
			return 0, err
		}
	}
	return amount, nil
}

//CheckForMail check if mail already exist
func CheckForMail(email string, db *sql.DB) bool {
	rows, err := db.Query("SELECT email FROM users")
	checkError(err)
	var (
		mail    string
		isExist bool
	)

	for rows.Next() {
		err = rows.Scan(&mail)
		checkError(err)
		if email == mail {
			isExist = true
			break
		} else {
			isExist = false
		}
	}
	return isExist
}

//Register Save New username , password and email of a new user
func Register(user string, passw string, email string, db *sql.DB) int {

	var rslt int

	isexist := CheckForUser(user, db)
	emailExist := CheckForMail(email, db)
	if isexist == 1 && emailExist {
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

//Login Check if username and password entred are the same as the ones on db
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
		}
	}

	rows.Close()
	return result
}

// GetStuff get user credentials from db
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
	}
	return i, u, p, c
}

// GetUID get user id
func GetUID(user string, db *sql.DB) int {

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

// Update update credentials
func Update(id int, db *sql.DB, args ...string) (bool, bool, bool) {

	defer db.Close()
	var (
		uOk bool
		pOk bool
		cOk bool
	)

	for i := range args {
		if id != 0 && args[i] != "" && i == 0 {
			stmt, err := db.Prepare("UPDATE passwords SET user=? WHERE pwid = ?")
			checkError(err)
			stmt.Exec(args[i], id)
			uOk = true

		} else if id != 0 && args[i] != "" && i == 1 {
			stmt, err := db.Prepare("UPDATE passwords SET passw=? WHERE pwid = ?")
			checkError(err)
			stmt.Exec(args[i], id)
			pOk = true

		} else if id != 0 && args[i] != "" && i == 2 {
			stmt, err := db.Prepare("UPDATE passwords SET category=? WHERE pwid = ?")
			checkError(err)
			stmt.Exec(args[i], id)
			cOk = true
		}
	}
	return uOk, pOk, cOk
}

//UpdatePw update user's login password
func UpdatePw(password, email string, db *sql.DB) bool {
	h := sha256.New()
	h.Write([]byte(password))
	ph := h.Sum(nil)
	phx := strings.ToLower(base32.HexEncoding.EncodeToString(ph))

	stmt, err := db.Prepare("UPDATE users SET password=? WHERE email = ?")
	checkError(err)
	_, err = stmt.Exec(phx, email)
	checkError(err)
	return true
}

// Save save creds
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

//Delete deletes credetials by providing password id or category
func Delete(pwid int, category string, db *sql.DB) (bool, error) {
	defer db.Close()
	if category == "" && pwid != 0 {
		stmt, err := db.Prepare("DELETE FROM passwords WHERE pwid = ?")
		stmt.Exec(pwid)
		if err != nil {
			return false, err
		}
		return true, nil
	} else if category != "" && pwid != 0 {
		stmt, err := db.Prepare("DELETE FROM passwords WHERE pwid = ? AND category = ?")
		stmt.Exec(pwid, category)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
