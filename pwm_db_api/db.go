package pwm_db_api

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

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
	_, err = c.conn.Exec(fmt.Sprintf("CREATE DATABASE %v", testDbName))
	if err != nil {
		return "", err
	}
	out, err := pq.ParseURL(basicURL)
	if err != nil {
		return "", err
	}
	host := strings.Trim(strings.Split(strings.Split(out, " ")[1], "=")[1], "'")
	testDbPath := fmt.Sprintf("postgres://postgres:%v/%v?sslmode=disable", host, testDbName)
	return testDbPath, nil
}

// Close closes the postgres db connection.
func (d Db) Close() error {
	return d.conn.Close()
}

// GetStoredServerKey get server key if it's available in the db.
func (d Db) GetStoredServerKey() (key string, err error) {
	err = d.conn.QueryRow(
		`
SELECT server_key
FROM server`).Scan(&key)
	return
}

// StoreServerKey
func (d Db) StoreServerKey(key string) error {
	result, err := d.conn.Exec(
		`
INSERT into server(server_key) values($1)
		`, key)
	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return ErrInsertion
	}

	return nil
}

type RegistrationConfig struct {
	Username string
	Password string
	Key      string
}

func (d Db) InsertNewUser(config RegistrationConfig) error {
	result, err := d.conn.Exec(
		`
INSERT into user_t(username, password, key) VALUES($1, $2, $3)
		`, config.Username, config.Password, config.Key)

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
