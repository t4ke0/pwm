package pwm_db_api

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var SchemaFile = "schema.sql"

var (
	ErrNoRows    = sql.ErrNoRows
	ErrInsertion = errors.New("couldn't insert value")
)

// Db
type Db struct {
	conn *sql.DB
}

// New get new instance of Db.
// by providing the url of the postgres database.
func New(url string) (Db, error) {
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return Db{}, err
	}
	return Db{conn}, nil
}

// CreateTestingDatabase
func CreateTestingDatabase(basicURL string) (string, error) {
	const testDbName = "unit_test"
	c, err := New(basicURL)
	if err != nil {
		return "", err
	}
	var testDBExist string
	if err := c.conn.QueryRow(`
SELECT datname
FROM pg_database
WHERE datname = $1`, testDbName).Scan(&testDBExist); err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if testDBExist == "" {
		_, err = c.conn.Exec(fmt.Sprintf("CREATE DATABASE %v", testDbName))
		if err != nil {
			return "", err
		}
	}
	var testDbPath string
	var prev, next rune
	for i, u := range basicURL {
		if i != 0 && i != len(basicURL)-1 {
			prev, next = rune(basicURL[i-1]), rune(basicURL[i+1])
			if u == '/' && prev != '/' && next != '/' {
				testDbPath = basicURL[:i] + "/" + testDbName + "?sslmode=disable"
			}
		}
	}
	return testDbPath, nil
}

// ClearTestTables remove testing tables.
func (d Db) ClearTestTables() error {
	_, err := d.conn.Exec(`
DELETE FROM sessions;
DELETE FROM passwords;
DELETE FROM user_t;
`)
	return err
}

// InitDB initializes the tables.
func (d Db) InitDB() error {
	data, err := os.ReadFile(SchemaFile)
	if err != nil {
		return err
	}
	_, err = d.conn.Exec(string(data))
	if err != nil {
		return err
	}
	return nil
}

// Close closes the postgres db connection.
func (d Db) Close() error {
	return d.conn.Close()
}

func (d Db) GetServerEncryptionKey() (string, error) {
	var key string
	err := d.conn.QueryRow(`
SELECT server_key FROM server_encryption_key
`).Scan(&key)
	if err != nil && err != ErrNoRows {
		return key, err
	}
	return key, nil
}

// StoreServerKey store the server encryption key.
func (d Db) StoreServerKey(key string) error {
	result, err := d.conn.Exec(
		`
INSERT into server_encryption_key(server_key) values($1)
		`, key)
	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return ErrInsertion
	}

	return nil
}

// GetAuthServerKey ...
func (d Db) GetAuthServerKey() (string, error) {
	var authSrvKey string
	err := d.conn.QueryRow("SELECT auth_server_key FROM server_auth_key").Scan(&authSrvKey)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return authSrvKey, nil
}

// StoreAuthServerKey
func (d Db) StoreAuthServerKey(key string) error {
	result, err := d.conn.Exec(
		`
INSERT INTO server_auth_key(auth_server_key)
VALUES($1)
`, key)
	if err != nil {
		return err
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return ErrInsertion
	}
	return nil
}

// RegistrationConfig registration configuration
type RegistrationConfig struct {
	Username string
	Password string
	Email    string
	Key      string
}

// InsertNewUser insert new user into user_t table.
func (d Db) InsertNewUser(config RegistrationConfig) error {
	result, err := d.conn.Exec(
		`
INSERT into user_t(username, password, email, key) VALUES($1, $2, $3, $4)
		`, config.Username, config.Password, config.Email, config.Key)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n != 1 {
		return ErrInsertion
	}
	return nil
}

// LoadUserKey
func (d Db) LoadUserKey(username string) (userkey string, err error) {
	err = d.conn.QueryRow(
		`
SELECT key FROM user_t WHERE username = $1
		`, username).Scan(&userkey)

	return
}

// UserExists
func (d Db) UserExists(username string) (bool, error) {
	var count int
	err := d.conn.QueryRow(
		`
SELECT COUNT(*) FROM user_t
WHERE username = $1
`, username).Scan(&count)
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}

// EmailExists
func (d Db) EmailExists(email string) (bool, error) {
	var storedEmail string
	var err = d.conn.QueryRow(`SELECT email from user_t where email = $1`, email).Scan(&storedEmail)
	if err != nil && err != ErrNoRows {
		return false, err
	}
	if storedEmail == "" {
		return false, nil
	}
	return true, nil
}

type UserInfo struct {
	ID  int
	Key string

	Password string
}

func (d Db) GetUserAuthInfo(username string) (UserInfo, error) {
	var info UserInfo
	err := d.conn.QueryRow(`
SELECT id, password, key
FROM user_t
WHERE username = $1`, username).Scan(&info.ID, &info.Password, &info.Key)
	if err != nil {
		return info, err
	}
	return info, nil
}

// InsertNewSession
func (d Db) InsertNewSession(sessionID, jwtToken string, userID int, createdAt time.Time) error {
	result, err := d.conn.Exec(`
INSERT INTO sessions(session_id, jwt_token, user_id, created_at, revoked)
VALUES($1, $2, $3, $4, false)`, sessionID, jwtToken, userID, createdAt)
	if err != nil {
		return err
	}
	n, _ := result.RowsAffected()
	if n == 0 {
		return ErrInsertion
	}
	return nil
}

// RevokeSession revoke a session given the jwt token and returns an error.
func (d Db) RevokeSession(jwtToken string) error {
	_, err := d.conn.Exec(
		`
UPDATE sessions SET revoked = true
WHERE jwt_token = $1
`, jwtToken)
	if err != nil {
		return err
	}
	return nil
}

// Passwords structure that holds the user password's fields.
type Passwords struct {
	ID                int
	EncryptedPassword string
	Username          string
	Category          string
	Site              string

	err error
}

// GetUserPasswords get user passwords returns a channel of Passwords type.
func (d Db) GetUserPasswords(userID int) (<-chan Passwords, error) {

	rows, err := d.conn.Query(
		`
SELECT id, password, username, category, site
FROM passwords
WHERE user_id = $1`, userID)

	if err != nil {
		return nil, err
	}

	out := make(chan Passwords)

	go func() {
		defer close(out)
		defer rows.Close()
		for rows.Next() {
			var pw Passwords
			if err := rows.Scan(&pw.ID, &pw.EncryptedPassword, &pw.Username,
				&pw.Category, &pw.Site); err != nil {
				pw.err = err
				out <- pw
				return
			}
			out <- pw
		}
	}()

	return out, nil
}

// StoreUserPassword stores user password into passwords table. accepts userID
// and encrypted user password and returns an error if exists.
func (d Db) StoreUserPassword(userID int, password Passwords) error {
	result, err := d.conn.Exec(
		`INSERT INTO passwords(user_id, password, username, category, site)
 			VALUES($1, $2, $3, $4, $5)`, userID, password.EncryptedPassword, password.Username,
		password.Category, password.Site)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return ErrInsertion
	}

	return nil
}

type ElementToUpdate string

const (
	Password ElementToUpdate = "password"
	Username                 = "username"
	Category                 = "category"
	Site                     = "site"
)

// UpdateUserPassword update password Element of the user.
func (d Db) UpdateUserPassword(userID, passwordID int, itemsToUpdate map[ElementToUpdate]string) error {

	var (
		query     string = "UPDATE passwords set("
		values    []interface{}
		valsQuery string = "VALUES("
	)
	count := 0
	for k, v := range itemsToUpdate {
		query += fmt.Sprintf("%s", k)
		valsQuery += fmt.Sprintf("$%d", count+1)
		if count != len(itemsToUpdate)-1 {
			query += ","
			valsQuery += ","
		} else {
			query += ")"
			valsQuery += ")"
		}
		values = append(values, v)
		count++
	}
	query = fmt.Sprintf("%s %s WHERE user_id = $%d AND id = $%d ", query, valsQuery, count+1, count+2)

	values = append(values, userID, passwordID)
	_, err := d.conn.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

// DeletePassword
func (d Db) DeletePassword(userID, passwordID int) (err error) {
	_, err = d.conn.Exec(`DELETE FROM passwords where id = $1 AND user_id = $2`,
		passwordID, userID)
	return
}
