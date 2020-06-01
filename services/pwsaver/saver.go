package pwsaver

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"

	"../../sqlite"
	"../pwencrypter"
	"../serverenc"
)

// AddCreds func saves user credentials to the Database
// Load User encryption key and encrypt passwords then add them
//TODO: accept username and pwshow.UserList as input then Loop through the list and add each row to the database .
func AddCreds(user string, password string, category string, Cuser string) bool {
	var isOk bool
	db := sqlite.InitDb()
	uid := sqlite.GetUID(Cuser, db)
	//	key := pwencrypter.LoadKey(Cuser) // user ioutil.ReadFile instead
	key, err := ioutil.ReadFile(path.Join(serverenc.KeysDir, Cuser+".key"))
	if err != nil {
		log.Fatal(err)
	}
	//Load server key
	serverK := pwencrypter.LoadKey("server")
	// Decrypt user key for encrypting his password
	decKey := serverenc.DecryptUserKey(key, serverK)
	//TODO: add loop here
	Epw := pwencrypter.Encrypt(password, decKey)
	hexenc := make([]byte, hex.EncodedLen(len(Epw)))
	hex.Encode(hexenc, Epw)
	if ok := sqlite.Save(user, string(hexenc), category, uid, db); ok {
		isOk = true
	} else {
		isOk = false
	}
	return isOk
}
