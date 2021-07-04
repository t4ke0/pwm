package pwm_db_api

import (
	"database/sql"
	"errors"

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