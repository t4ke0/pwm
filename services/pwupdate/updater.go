package pwupdate

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"

	"../../sqlite"
	"../pwencrypter"
	"../serverenc"
)

// UpdateCreds Update Credentials
func UpdateCreds(id int, username string, user string, password string, category string) map[string]bool {
	db := sqlite.InitDb()
	//uid := sqlite.GetUID(username, db)
	ukey, err := ioutil.ReadFile(path.Join(serverenc.KeysDir, username+".key"))
	if err != nil {
		log.Fatal(err)
	}
	// load server key
	serverk := pwencrypter.LoadKey("server")
	// Decrypt user key
	decK := serverenc.DecryptUserKey(ukey, serverk)
	encp := pwencrypter.Encrypt(password, decK)
	hexenc := make([]byte, hex.EncodedLen(len(encp)))
	hex.Encode(hexenc, encp)
	ud := sqlite.Update(id, db, user, string(hexenc), category)
	m := make(map[string]bool)

	for i, j := range ud {
		if i == 0 && j == 1 {
			m["user"] = true
		} else if i == 0 && j == 0 {
			m["user"] = false
		}
		if i == 1 && j == 1 {
			m["password"] = true
		} else if i == 1 && j == 0 {
			m["password"] = false
		}
		if i == 2 && j == 1 {
			m["category"] = true
		} else if i == 2 && j == 0 {
			m["category"] = false
		}
	}
	return m
}
