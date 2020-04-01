package handler

import (
	"fmt"
	"net/http"
	"strings"

	"../dialer"
)

// ReqHandler handles http request made by the user
func ReqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
	reqPath := strings.Trim(r.URL.Path, "/")
	switch reqPath {
	case "":
		dialer.ServeHome(w, r)
	case "login":
		dialer.ServeLogin(w, r)
	case "register":
		dialer.ServeRegister(w, r)
	case "add":
		dialer.ServeAdd(w, r)
	case "logout":
		dialer.HandleLogout(w, r)
	case "show":
		dialer.ServeShow(w, r)
	case "update":
		dialer.ServeUpdate(w, r)
	case "delete":
		dialer.ServeDelete(w, r)
	default:
		fmt.Fprintf(w, "404 Not Found", http.StatusNotFound)
	}
}
