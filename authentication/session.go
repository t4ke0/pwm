package authentication

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

const sessionCookie = "session"

func SetSession(username string, w http.ResponseWriter) error {
	value := map[string]string{
		"name": username,
	}
	var err error
	encoded, err := cookieHandler.Encode(sessionCookie, value)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:   sessionCookie,
		Value:  encoded,
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(w, cookie)
	return nil
}

func GetUsername(r *http.Request) (username string) {
	if cookie, err := r.Cookie(sessionCookie); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			username = cookieValue["name"]
		}
	}
	return username
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   sessionCookie,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func CheckCookie(r *http.Request) bool {
	var ok bool
	if cookie, _ := r.Cookie(sessionCookie); cookie != nil {
		ok = true
	} else {
		ok = false
	}
	return ok
}
