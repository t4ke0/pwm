package dialer

//TODO: don't send cookiename and it value , get the name and value from response header
//TODO: add a go routine to send emails

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../../authentication"
	"../../identityprovider"
	"../../services/emailsender"
	"../../services/pwdelete"
	"../../services/pwsaver"
	"../../services/pwshow"
)

// User Struct
type User struct {
	IsEmpty  bool            `json:"isEmpty"`
	CredList pwshow.UserList `json:"CredList"`
}

// CookieUser struct holds username of the current user in the session
type CookieUser struct {
	Username string `json:Username`
}

//Register IsReg value is sent to inform that the user is successfully registred
type Register struct {
	IsReg bool `json:"IsReg"`
}

//Creds send Updated if the db is updated
type Creds struct {
	Updated bool `json:"Updated"`
}

//Login login struct holds values that we need in the frontend to check if we have successfully logged in
type Login struct {
	IsLog       bool   `json:"IsLog"`
	CookieName  string `json:"CookieName"`
	CookieValue string `json:"CookieValue"`
}

//Token is the code that we send to the user to restore his password
type Token struct {
	Code string
}

//Password password struct holds a value which is Updated to inform the frontend that the user password is updated
type Password struct {
	Updated bool `json:"Updated"`
}

//Email has Response and IsEqual instances we need them in Password recovery process.
type Email struct {
	Response bool `json:"Response"`
	IsEqual  bool `json:"IsEqual"`
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

//CookieDecode decode the cookie and get the username
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
	if f := HandlePost(r); len(f) != 0 && user != "" {
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

var t Token

//ServepwForget function that handles password recovery process
func ServepwForget(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	e := &Email{}
	if f := HandlePost(r); len(f) != 0 {
		password := strings.Join(f["npassword"], "")
		if email := strings.Join(f["email"], ""); email != "" && password == "" {
			if mailExist := authentication.CheckMail(email); mailExist {
				rand.Seed(time.Now().UnixNano())
				gencode := rand.Perm(6)
				var c string
				// Convert []int into string
				for _, n := range gencode {
					c += strconv.Itoa(n)
				}
				if sent, err := emailsender.SendCode(c, email); err != nil {
					log.Fatal(err)
				} else if sent {
					t.Code = c
					e.Response = true
					json.NewEncoder(w).Encode(e)
				}
			} else {
				t.Code = ""
				e.Response = false
				json.NewEncoder(w).Encode(e)
			}
		} else if code := strings.Join(f["code"], ""); code != "" {
			if code == t.Code {
				e.IsEqual = true
				json.NewEncoder(w).Encode(e)
			} else {
				e.IsEqual = false
				json.NewEncoder(w).Encode(e)
			}
		} else if password != "" && email != "" {
			p := &Password{}
			isUpdated := authentication.UpdatePassword(email, password)
			if isUpdated {
				p.Updated = isUpdated
				json.NewEncoder(w).Encode(p)
			} else {
				p.Updated = false
				json.NewEncoder(w).Encode(p)
			}
		}
	}
}

var d pwshow.UserList

func handleJSONBody(r *http.Request) error {
	if r.Body != nil {
		err := json.NewDecoder(r.Body).Decode(&d)
		return err
	}
	return nil
}

//ServeCreds update and delete User creds if the front-end credential has been changed
func ServeCreds(w http.ResponseWriter, r *http.Request) {
	handleOption(w, r)
	username := authentication.GetUsername(r)
	err := handleJSONBody(r)
	CheckError(err)
	if username != "" && d != nil {
		// 1st delete user creds
		isDeleted := pwdelete.DeleteCreds(username)
		if isDeleted {
			//Now we should update user creds Here
			for _, n := range d {
				pwsaver.AddCreds(n.Username, n.Password, n.Category, username)
			}
			c := &Creds{Updated: true}
			json.NewEncoder(w).Encode(c)
		}
	}
}
