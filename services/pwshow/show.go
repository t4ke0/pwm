package pwshow

import (
	"../../sqlite"
	"../pwencrypter"
	"../serverenc"
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"
)

type UserStuff struct {
	Id       string
	Username string
	Password string
	Category string
}

type UserList []UserStuff

var (
	FinalList UserList
	U         = UserStuff{}
)

// Add Creds To the FinalList
// Load User Encryption Key and decrypt passwrds
func AddToList(i []string, u []string, p []string, c []string, user string) UserList {
	//Clear the FinalList Each Call
	FinalList = FinalList[:0]
	//	key := pwencrypter.LoadKey(user)
	// Loading User Key
	userk, err := ioutil.ReadFile(path.Join(serverenc.KeysDir, user+".key"))
	if err != nil {
		log.Fatal(err)
	}
	// Load Server Key
	serverK := pwencrypter.LoadKey("server")
	// Decrypt User Key
	decKey := serverenc.DecryptUserKey(userk, serverK)
	for x, _ := range i {
		src := []byte(p[x])
		dst := make([]byte, hex.DecodedLen(len(src)))
		n, err := hex.Decode(dst, src)
		if err != nil {
			log.Fatal(err)
		}
		U.Id, U.Username, U.Password, U.Category = i[x], u[x], pwencrypter.Decrypt(dst[:n], decKey), c[x]
		FinalList = append(FinalList, U)
	}
	return FinalList
}

// ShowCreds Function Get user Stuff From the database
func ShowCreds(user string, category string) UserList {
	var Fl UserList
	db := sqlite.InitDb()
	uid := sqlite.GetUid(user, db)
	if category != "" {
		i, u, p, c := sqlite.GetStuff(uid, category, db)
		Fl = AddToList(i, u, p, c, user)
	} else {
		i, u, p, c := sqlite.GetStuff(uid, category, db)
		Fl = AddToList(i, u, p, c, user)
	}
	return Fl
}
