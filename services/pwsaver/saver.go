package pwsaver

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"

	"../../sqlite"
	"../pwencrypter"
	"../pwshow"
	"../serverenc"
)

// AddCreds func saves user credentials to the Database
// Load User encryption key and encrypt passwords then add them
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

//ParseAndAdd loop through UserList type and save credentials
func ParseAndAdd(credList pwshow.UserList, username string) bool {
	var svOk bool
	for _, n := range credList {
		svOk = AddCreds(n.Username, n.Password, n.Category, username)
	}
	return svOk
}
