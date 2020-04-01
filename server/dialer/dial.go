package dialer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../../authentication"
	"../../identityprovider"
	"../../services/pwdelete"
	"../../services/pwsaver"
	"../../services/pwshow"
	"../../services/pwupdate"
)

// Static indicates Static Files "html files"
const Static string = "./server/staticfiles/"

// User Struct
type User struct {
	Username  string
	Ok        bool
	Cookie    bool
	IsEmpty   bool
	Updated   bool
	UserExist bool
	CredList  pwshow.UserList
}

type usr User

var u usr

// HandleGet function handles GET http method
// By Serving the html files and passing User Struct fields into templates
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

// HandlePost function handles Post http method
// Gets input from form and returns it as map[string][]string
func HandlePost(r *http.Request) map[string][]string {
	var m map[string][]string

	if r.Method == "POST" {
		r.ParseForm()
		m = r.Form
	}
	return m
}

// CheckError function check if err not nil then log the error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ServeHome serve the HOME Page for the user
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

// ServeRegister serve the register Page for the user
func ServeRegister(w http.ResponseWriter, r *http.Request) {
	cookie := authentication.CheckCookie(r)
	u.Cookie = cookie
	HandleGet(w, r, "register.html", u)
	if f := HandlePost(r); len(f) != 0 {
		user := strings.Join(f["user"], "")
		password := strings.Join(f["passw"], "")
		email := strings.Join(f["email"], "")

		if ok := identityprovider.GetRegister(r, user, password, email); ok {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			u.UserExist = true
			http.Redirect(w, r, "/register", http.StatusFound)
		}
	}
}

// ServeLogin serves login page for the user
func ServeLogin(w http.ResponseWriter, r *http.Request) {
	if cookie := authentication.CheckCookie(r); !cookie {
		HandleGet(w, r, "login.html")
	} else {
		fmt.Fprintf(w, "You are Already logged in")
	}

	if f := HandlePost(r); len(f) != 0 {
		user := strings.Join(f["user"], "")
		password := strings.Join(f["passw"], "")
		if ok := identityprovider.GetLoggedin(w, r, user, password); ok {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Fprintf(w, "Wrong username or password!")
		}
	}
}

// HandleLogout Logout the user
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	if cookie := authentication.CheckCookie(r); cookie {
		authentication.ClearSession(w)
		fmt.Fprintf(w, "You were Logged out")
	} else {
		fmt.Fprintf(w, "You are already logged out")
	}
}

// ServeShow Show User credentials
func ServeShow(w http.ResponseWriter, r *http.Request) {
	var l pwshow.UserList
	if cookie := authentication.CheckCookie(r); cookie {
		HandleGet(w, r, "show.html", u)
		user := authentication.GetUsername(r)
		if f := HandlePost(r); len(f) != 0 {
			category := strings.Join(f["category"], "")
			l = pwshow.ShowCreds(user, category)
			if len(l) != 0 {
				u.CredList = l
				http.Redirect(w, r, "/show", http.StatusFound)
			} else {
				u.IsEmpty = true
				http.Redirect(w, r, "/show", http.StatusFound)
			}
		}
	} else {
		fmt.Fprintf(w, "You Are Not Loggedin")
	}
}

// ServeUpdate Get item that Should be updated then sent them to Update service
func ServeUpdate(w http.ResponseWriter, r *http.Request) {
	if cookie := authentication.CheckCookie(r); cookie {
		username := authentication.GetUsername(r)
		HandleGet(w, r, "update.html", u)
		u.Updated = false
		if f := HandlePost(r); len(f) != 0 {
			id := strings.Join(f["id"], "")
			user := strings.Join(f["user"], "")
			password := strings.Join(f["passw"], "")
			category := strings.Join(f["catg"], "")
			iid, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			m := pwupdate.UpdateCreds(iid, username, user, password, category)
			if m["user"] == true {
				u.Updated = true
			} else if m["password"] == true {
				u.Updated = true
			} else if m["category"] == true {
				u.Updated = true
			}
			http.Redirect(w, r, "/update", http.StatusFound)
		}
	} else {
		fmt.Fprintf(w, "You Are Not Loggedin")
	}
}

// ServeAdd Get items that should be added and send them to the save credential service to save them
func ServeAdd(w http.ResponseWriter, r *http.Request) {
	if cookie := authentication.CheckCookie(r); cookie {
		u.Cookie = true
		Tuser := authentication.GetUsername(r)
		u.Username = Tuser
		// Serving the html page
		HandleGet(w, r, "add.html", u)
		u.Ok = false
		// handling the post request
		if f := HandlePost(r); len(f) != 0 {
			username := strings.Join(f["user"], "")
			passw := strings.Join(f["passw"], "")
			category := strings.Join(f["category"], "")
			if IsSaved := pwsaver.AddCreds(username, passw, category, Tuser); IsSaved {
				u.Ok = true
				http.Redirect(w, r, "/add", http.StatusFound)
			}
		}
	} else {
		fmt.Fprintf(w, "You Are Not Loggedin")
	}
}

// ServeDelete Get the id number of the cred that should be deleted and send them to the delete service
func ServeDelete(w http.ResponseWriter, r *http.Request) {
	if cookie := authentication.CheckCookie(r); cookie {
		HandleGet(w, r, "delete.html", u)
		u.Ok = false
		if f := HandlePost(r); len(f) != 0 {
			id, err := strconv.Atoi(strings.Join(f["id"], ""))
			CheckError(err)
			if isDeleted := pwdelete.DeleteCreds(id); isDeleted {
				u.Ok = true
				http.Redirect(w, r, "/delete", http.StatusFound)
			}
		}
	} else {
		fmt.Fprintf(w, "You are not loggeding")
	}
}
