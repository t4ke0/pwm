package identityprovider

import (
	"../authentication"
	"../services/pwencrypter"
	"../services/serverenc"
	"log"
	"net/http"
)

func GetLoggedin(w http.ResponseWriter, r *http.Request, user string, password string) bool {
	var rr bool
	if ok := authentication.CheckCookie(r); ok {
		rr = false
	} else {
		if logged := authentication.Login(user, password); logged {
			if err := authentication.SetSession(user, w); err != nil {
				log.Fatal(err)
			}
			rr = true
		} else {
			rr = false
		}
	}
	return rr
}

func GetRegister(r *http.Request, user string, password string, email string) bool {
	var rr bool
	if ok := authentication.CheckCookie(r); ok {
		rr = false
	} else {
		authentication.Register(user, password, email)
		// Generating the encryption key for the user then save it in the keys directory
		key := pwencrypter.GenKeyP(password)
		// Load server key and Encrypt user key
		srvk := pwencrypter.LoadKey("server")
		nuserK := serverenc.EncryptUserKey(key, srvk)
		if isSaved := pwencrypter.SaveKey(nuserK, user); isSaved {
			rr = true
		}
	}
	return rr
}

func GetLoggedout(w http.ResponseWriter, r *http.Request) bool {
	var rr bool
	if ok := authentication.CheckCookie(r); ok {
		authentication.ClearSession(w)
		rr = true
	} else {
		rr = false
	}
	return rr
}
