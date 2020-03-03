package dialer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const static string = "./staticfiles/"

// temp struct
type USER struct {
	Username string
}

type usr USER

func MethodCheck(w http.ResponseWriter, r *http.Request, tempt string, args ...usr) map[string][]string {
	var m map[string][]string

	switch req := r.Method; req {
	case "GET":
		t, err := template.ParseFiles(static + tempt)
		CheckError(err)
		if len(args) < 1 {
			t.Execute(w, nil)
		} else {
			t.Execute(w, args[0])
		}
	case "POST":
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
		t, err := template.ParseFiles(static + "home.html")
		CheckError(err)
		t.Execute(w, nil)
	case "POST":
		fmt.Fprintf(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func ServeRegister(w http.ResponseWriter, r *http.Request) {

	f := MethodCheck(w, r, "register.html")
	fmt.Println(strings.Join(f["user"], ""))

}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	MethodCheck(w, r, "login.html")
	//TODO get form values
}

func ServeShow(w http.ResponseWriter, r *http.Request) {
	MethodCheck(w, r, "show.html")
	//TODO ....
}

func ServeUpdate(w http.ResponseWriter, r *http.Request) {
	MethodCheck(w, r, "update.html")
	//TODO ....
}

func ServeAdd(w http.ResponseWriter, r *http.Request) {
	MethodCheck(w, r, "add.html")
}
func ServeDelete(w http.ResponseWriter, r *http.Request) {
	MethodCheck(w, r, "delete.html")
}
