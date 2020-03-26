package pwupdate

import (
	"encoding/hex"

	"../../sqlite"
	"../pwencrypter"
)

func UpdateCreds(id int, username string, user string, password string, category string) map[string]bool {
	db := sqlite.InitDb()
	//uid := sqlite.GetUid(username, db)
	key := pwencrypter.LoadKey(username)
	encp := pwencrypter.Encrypt(password, key)
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
