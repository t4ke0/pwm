package identityprovider

import (
	"github.com/TaKeO90/pwm/authentication"
	"log"
	"net/http"
)

//TODO use Login and Register Functions on authentication package
//TODO sqlite API then try if everything is OK

func GetLoggedin(w http.ResponseWriter, r *http.Request, user string, password string) int {
	var rr int
	if ok := authentication.CheckCookie(r); ok {
		rr = 0
	} else {
		if logged := authentication.Login(user, password); logged {
			err := authentication.SetSession(user, w)
			if err != nil {
				log.Fatal(err)
			}
			rr = 1
		} else {
			rr = 0
		}
	}
	return rr
}

func GetRegister(r *http.Request, user string, password string, email string) int {
	var rr int
	if ok := authentication.CheckCookie(r); ok {
		rr = 0
	} else {
		authentication.Register(user, password, email)
		rr = 1
	}
	return rr
}

func GetLoggedout(w http.ResponseWriter, r *http.Request) int {
	var rr int
	if ok := authentication.CheckCookie(r); ok {
		authentication.ClearSession(w)
		rr = 1
	} else {
		rr = 0
	}
	return rr

}
