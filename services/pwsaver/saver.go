package pwsaver

import (
	"../../sqlite"
	"../pwencrypter"
	"encoding/hex"
)

//TODO Encrypt Users passwords With a symmetric encryption and store all users keys in a folder under naem keys then encrypt this folder with another symetric key for the server

// AddCreds func saves user credentials to the Database
// Load User encryption key and encrpyt passwords then add them
func AddCreds(user string, password string, category string, Cuser string) bool {
	var isOk bool
	db := sqlite.InitDb()
	uid := sqlite.GetUid(Cuser, db)
	key := pwencrypter.LoadKey(Cuser)
	Epw := pwencrypter.Encrypt(password, key)
	hexenc := make([]byte, hex.EncodedLen(len(Epw)))
	hex.Encode(hexenc, Epw)
	if ok := sqlite.Save(user, string(hexenc), category, uid, db); ok {
		isOk = true
	} else {
		isOk = false
	}
	return isOk
}
