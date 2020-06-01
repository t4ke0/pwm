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

// UserStuff a struct which it field holds user credential fields
type UserStuff struct {
	Username string
	Password string
	Category string
}

//UserList list of UserStuff type
type UserList []UserStuff

var (
	//FinalList variable where we store user creds
	FinalList UserList
	//U variable where we put user creds before we append them to the FinalList
	U = UserStuff{}
)

//AddToList Add Creds To the FinalList
// Load User Encryption Key and decrypt passwrds
func AddToList(u []string, p []string, c []string, user string) UserList {
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
	for x := range u {
		src := []byte(p[x])
		dst := make([]byte, hex.DecodedLen(len(src)))
		n, err := hex.Decode(dst, src)
		if err != nil {
			log.Fatal(err)
		}
		U.Username, U.Password, U.Category = u[x], pwencrypter.Decrypt(dst[:n], decKey), c[x]
		FinalList = append(FinalList, U)
	}
	return FinalList
}

// ShowCreds Function Get user Stuff From the database
func ShowCreds(user string, category string) UserList {
	var Fl UserList
	db := sqlite.InitDb()
	uid := sqlite.GetUID(user, db)
	if category != "" {
		_, u, p, c := sqlite.GetStuff(uid, category, db)
		Fl = AddToList(u, p, c, user)
	} else {
		_, u, p, c := sqlite.GetStuff(uid, category, db)
		Fl = AddToList(u, p, c, user)
	}
	return Fl
}
