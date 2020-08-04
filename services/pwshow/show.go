package pwshow

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"
	"strconv"

	"github.com/TaKeO90/pwm/services/pwencrypter"
	"github.com/TaKeO90/pwm/services/serverenc"
	"github.com/TaKeO90/pwm/sqlite"
)

// UserStuff a struct which it field holds user credential fields
type UserStuff struct {
	Username string
	Password string
	Category string
}

//CompareS struct who has ID as pwid of credentials and credentials it selfs
type CompareS struct {
	ID    int
	Creds UserStuff
}

//UserList list of UserStuff type
type UserList []UserStuff

var (
	//FinalList variable where we store user creds
	FinalList   UserList
	compareList []CompareS
	//U variable where we put user creds before we append them to the FinalList
	U = UserStuff{}
)

// AddToList Add Creds To the FinalList
// Load User Encryption Key and decrypt passwrds
func AddToList(i, u, p, c []string, user string) UserList {
	//Clear the FinalList Each Call
	FinalList = FinalList[:0]
	compareList = compareList[:0]
	cU := CompareS{}
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
		id, _ := strconv.Atoi(i[x])
		cU.ID, cU.Creds.Username, cU.Creds.Password, cU.Creds.Category = id, u[x], pwencrypter.Decrypt(dst[:n], decKey), c[x]
		FinalList = append(FinalList, U)
		compareList = append(compareList, cU)
	}
	return FinalList
}

// GetPWID get password ID
func GetPWID(user, pw, catg string) int {
	if len(compareList) != 0 {
		for _, n := range compareList {
			if n.Creds.Username == user && n.Creds.Password == pw && n.Creds.Category == catg {
				return n.ID
			}
		}
	}
	return 0
}

// ShowCreds Function Get user Stuff From the database
func ShowCreds(user string, category string) UserList {
	var Fl UserList
	db := sqlite.InitDb()
	uid := sqlite.GetUID(user, db)
	if category != "" {
		i, u, p, c := sqlite.GetStuff(uid, category, db)
		Fl = AddToList(i, u, p, c, user)
	} else {
		i, u, p, c := sqlite.GetStuff(uid, category, db)
		Fl = AddToList(i, u, p, c, user)
	}
	return Fl
}
