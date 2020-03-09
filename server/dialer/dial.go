package dialer

import (
	"../../authentication"
	"../../identityprovider"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const Static string = "./server/staticfiles/"

type User struct {
	Username string
	Ok       bool
	Cookie   bool
}

type usr User

var u usr

func HandleGet(w http.ResponseWriter, r *http.Request, tempt string, args ...usr) {

	if req := r.Method; req == "GET" {
		t, err := template.ParseFiles(Static + tempt)
		CheckError(err)
		if len(args) < 1 {
			t.Execute(w, nil)
		} else {
			t.Execute(w, args[0])

		}
	}
}

func HandlePost(r *http.Request) map[string][]string {
	var m map[string][]string

	if r.Method == "POST" {
		r.ParseForm()
		m = r.Form
	}
	return m
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		user := authentication.GetUsername(r)
		p := User{Username: user}
		t, err := template.ParseFiles(Static + "home.html")
		CheckError(err)
		t.Execute(w, p)

	case "POST":
		fmt.Fprintf(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func ServeRegister(w http.ResponseWriter, r *http.Request) {
	cookie := authentication.CheckCookie(r)
	u.Cookie = cookie
	HandleGet(w, r, "register.html", u)
	if f := HandlePost(r); len(f) != 0 {
		user := strings.Join(f["user"], "")
		password := strings.Join(f["passw"], "")
		email := strings.Join(f["email"], "")

		if ok := identityprovider.GetRegister(r, user, password, email); ok == 1 {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			http.Redirect(w, r, "/register", http.StatusFound)
		}
	}

}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	if cookie := authentication.CheckCookie(r); !cookie {
		HandleGet(w, r, "login.html")
	} else {
		fmt.Fprintf(w, "You are Already logged in")
	}

	if f := HandlePost(r); len(f) != 0 {
		user := strings.Join(f["user"], "")
		password := strings.Join(f["passw"], "")
		if ok := identityprovider.GetLoggedin(w, r, user, password); ok == 1 {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Fprintf(w, "Wrong username or password!")
		}
	}
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	if cookie := authentication.CheckCookie(r); cookie {
		authentication.ClearSession(w)
		fmt.Fprintf(w, "You were Logged out")
	} else {
		fmt.Fprintf(w, "You are already logged out")
	}
}

func ServeShow(w http.ResponseWriter, r *http.Request) {
	HandleGet(w, r, "show.html")
	//TODO ....
}

func ServeUpdate(w http.ResponseWriter, r *http.Request) {
	HandleGet(w, r, "update.html")
	//TODO ....
}

func ServeAdd(w http.ResponseWriter, r *http.Request) {
	HandleGet(w, r, "add.html")
}
func ServeDelete(w http.ResponseWriter, r *http.Request) {
	HandleGet(w, r, "delete.html")
}
