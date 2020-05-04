package dialer

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"../../authentication"
	"../../identityprovider"
	"../../services/emailsender"
	"../../services/pwdelete"
	"../../services/pwsaver"
	"../../services/pwshow"
	"../../services/pwupdate"
)

// User Struct
type User struct {
	Username string          `json:"Username"`
	Ok       bool            `json:"Ok"`
	IsEmpty  bool            `json:"isEmpty"`
	Updated  bool            `json:"Updated"`
	CredList pwshow.UserList `json:"CredList"`
}

// CookieUser
type CookieUser struct {
	Username string `json:Username`
}

//Register
type Register struct {
	IsReg bool `json:"IsReg"`
}

//Login
type Login struct {
	IsLog       bool   `json:"IsLog"`
	CookieName  string `json:"CookieName"`
	CookieValue string `json:"CookieValue"`
}

type Token struct {
	Code string `json:"Code"`
}

// HandlePost function handles Post http method
// Gets input from form and returns it as map[string][]string
func HandlePost(r *http.Request) map[string][]string {
	var m map[string][]string
	if r.Method == "POST" {
		r.ParseMultipartForm(0)
		//r.ParseForm()
		m = r.Form
	}
	return m
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST,GET")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Expose-Headers", "Set-Cookie") // Not Working TOBE removed
	(*w).Header().Set("Content-Type", "application/json")
}

func handleOption(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
}

// CheckError function check if err not nil then log the error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ServeRegister Serve Registration
func ServeRegister(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	if f := HandlePost(r); len(f) != 0 {
		user := strings.Join(f["user"], "")
		password := strings.Join(f["passw"], "")
		email := strings.Join(f["email"], "")
		regtr := &Register{}
		if ok := identityprovider.GetRegister(r, user, password, email); ok {
			regtr.IsReg = true
			json.NewEncoder(w).Encode(regtr)
		} else {
			regtr.IsReg = false
			json.NewEncoder(w).Encode(regtr)
		}
	}
}

// ServeLogin handle login process
func ServeLogin(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	if f := HandlePost(r); len(f) != 0 {
		user := strings.Join(f["user"], "")
		password := strings.Join(f["passw"], "")
		logg := &Login{}
		if cookie, ok := identityprovider.GetLoggedin(w, r, user, password); ok {
			logg.CookieName = cookie.Name
			logg.CookieValue = cookie.Value
			logg.IsLog = true
			json.NewEncoder(w).Encode(logg)
		} else {
			logg.IsLog = false
			json.NewEncoder(w).Encode(logg)
		}
	}
}

//CookieDecode
func CookieDecode(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	uname := authentication.GetUsername(r)
	c := &CookieUser{uname}
	json.NewEncoder(w).Encode(c)
}

// ServeShow Show User credentials
func ServeShow(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	var l pwshow.UserList
	user := authentication.GetUsername(r)
	if f := HandlePost(r); len(f) != 0 {
		category := strings.Join(f["category"], "")
		l = pwshow.ShowCreds(user, category)
		u := &User{}
		if len(l) != 0 {
			u.CredList = l
			json.NewEncoder(w).Encode(u)
		} else {
			u.IsEmpty = true
			json.NewEncoder(w).Encode(u)
		}
	}
}

// ServeUpdate Get item that Should be updated then sent them to Update service
func ServeUpdate(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	u := &User{}
	if cookie := authentication.CheckCookie(r); cookie {
		username := authentication.GetUsername(r)
		u.Updated = false
		if f := HandlePost(r); len(f) != 0 {
			id := strings.Join(f["id"], "")
			user := strings.Join(f["user"], "")
			password := strings.Join(f["passw"], "")
			category := strings.Join(f["catg"], "")
			iid, err := strconv.Atoi(id)
			CheckError(err)
			m := pwupdate.UpdateCreds(iid, username, user, password, category)
			if m["user"] == true {
				u.Updated = true
			} else if m["password"] == true {
				u.Updated = true
			} else if m["category"] == true {
				u.Updated = true
			}
			json.NewEncoder(w).Encode(u)
		}
	} else {
		u.Updated = false
		json.NewEncoder(w).Encode(u)
	}
}

// ServeAdd Get items that should be added and send them to the save credential service to save them
func ServeAdd(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	u := &User{}
	if cookie := authentication.CheckCookie(r); cookie {
		Tuser := authentication.GetUsername(r)
		u.Username = Tuser
		u.Ok = false
		// handling the post request
		if f := HandlePost(r); len(f) != 0 {
			username := strings.Join(f["user"], "")
			passw := strings.Join(f["passw"], "")
			category := strings.Join(f["category"], "")
			if IsSaved := pwsaver.AddCreds(username, passw, category, Tuser); IsSaved {
				u.Ok = true
				json.NewEncoder(w).Encode(u)
			}
		}
	} else {
		u.Ok = false
		json.NewEncoder(w).Encode(u)
	}
}

// ServeDelete Get the id number of the cred that should be deleted and send them to the delete service
func ServeDelete(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	u := &User{}
	if cookie := authentication.CheckCookie(r); cookie {
		u.Ok = false
		if f := HandlePost(r); len(f) != 0 {
			id, err := strconv.Atoi(strings.Join(f["id"], ""))
			CheckError(err)
			if isDeleted := pwdelete.DeleteCreds(id); isDeleted {
				u.Ok = true
				json.NewEncoder(w).Encode(u)
			}
		}
	} else {
		u.Ok = false
		json.NewEncoder(w).Encode(u)
	}
}

func ServepwForget(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	t := &Token{}
	if f := HandlePost(r); len(f) != 0 {
		if email := strings.Join(f["email"], ""); email != "" {
			if mailExist := authentication.CheckMail(email); mailExist {
				gencode := rand.Perm(6)
				var c string
				// Conver []int into string
				for _, n := range gencode {
					c += strconv.Itoa(n)
				}
				//TODO: send generated code vim email "c"
				if sent, err := emailsender.SendCode(c, email); err != nil {
					log.Fatal(err)
				} else if sent {
					t.Code = c
					json.NewEncoder(w).Encode(t)
				}
			} else {
				t.Code = ""
				json.NewEncoder(w).Encode(t)
			}
		}
	}
}
