package pwshow

import (
	"../../sqlite"
	"../pwencrypter"
	"../serverenc"
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"
	"strconv"
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

//GetPWID get password ID
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

//Compare compare front-end data and  back-end's
func Compare(Flist UserList, category string) (map[UserStuff]UserStuff, UserList, UserList, bool) {
	m := make(map[UserStuff]UserStuff)
	var save UserList
	var delt UserList
	var ctg bool
	/* if length of the data in front-end equal to the data in backend */
	if len(FinalList) == len(Flist) {
		mm, _, isctg := oldVsNew(FinalList, Flist, category, 0)
		m = mm
		ctg = isctg

		/* if length of front-end data is bigger than the data in the backend */
	} else if len(FinalList) < len(Flist) {
		Fl := Flist[:len(FinalList)]
		mm, _, isctg := oldVsNew(FinalList, Fl, category, 0)
		if mm != nil {
			m = mm
			ctg = isctg
		}
		_, add, isctg := oldVsNew(FinalList, Flist, category, 1)
		if add != nil {
			save = add
			ctg = isctg
		}
		/* if length of front-end's data is less than the backend's */
	} else if len(FinalList) > len(Flist) {
		Fl := make(UserList, len(FinalList))
		Fl = clone(Flist, Fl)
		_, deleted, isctg := oldVsNew(FinalList, Fl, category, 2)
		delt = deleted
		ctg = isctg
	}
	return m, save, delt, ctg
}

func clone(arr1, arr2 UserList) UserList {
	for i := range arr1 {
		arr2[i] = arr1[i]
	}
	return arr2
}

func oldVsNew(old, New UserList, category string, mode int) (map[UserStuff]UserStuff, UserList, bool) {
	m := make(map[UserStuff]UserStuff)
	var additional UserList
	var Todelete UserList
	var isCtg bool

	if category == "" && mode == 0 {
		for ind, n := range old {
			if n.Username != New[ind].Username || n.Password != New[ind].Password || n.Category != New[ind].Category {
				m[n] = New[ind]
			}
		}
		return m, nil, false
	} else if category != "" && mode == 0 {
		for ind, n := range old {
			if n.Username != New[ind].Username || n.Password != New[ind].Password && n.Category == category {
				m[n] = New[ind]
				isCtg = true
			}
		}
		return m, nil, isCtg
	} else if mode == 1 {
		lstElms := make(UserList, (len(New) - len(old)))
		lstElms1 := clone(New[len(old):], lstElms)

		for inx, i := range old {
			if inx == len(lstElms1) {
				break
			}
			if category == "" {
				if i.Username != lstElms1[inx].Username || i.Password != lstElms1[inx].Password || i.Category != lstElms1[inx].Category {
					additional = append(additional, lstElms1[inx])
				}
			} else {
				if i.Username != lstElms1[inx].Username || i.Password != lstElms1[inx].Password || i.Category == category {
					additional = append(additional, lstElms1[inx])
					isCtg = true
				}
			}
		}
		if len(lstElms1) > len(old) {
			lstElms1 = lstElms1[len(old):]
			if len(lstElms1) < len(old) {
				lst := make(UserList, len(old))
				lst1 := clone(lstElms1, lst)
				if category != "" {
					for ind, i := range FinalList {
						if i.Username != lst1[ind].Username || i.Password != lst1[ind].Password && i.Category == category {
							if lst1[ind].Username != "" && lst[ind].Password != "" && lst[ind].Category != "" {
								additional = append(additional, lst1[ind])
								isCtg = true
							}
						}
					}
				} else {
					for ind, i := range FinalList {
						if i.Username != lst1[ind].Username || i.Password != lst1[ind].Password && i.Category != lst1[ind].Category {
							if lst1[ind].Username != "" && lst[ind].Password != "" && lst[ind].Category != "" {
								additional = append(additional, lst1[ind])
							}
						}
					}
				}
			}
		}
		return nil, additional, isCtg
	} else if mode == 2 {
		if category == "" {
			for index, n := range old {
				if n.Username != New[index].Username || n.Password != New[index].Password || n.Category != New[index].Category {
					Todelete = append(Todelete, n)
				}
			}
		} else {
			for index, n := range old {
				if n.Username != New[index].Username || n.Password != New[index].Password && n.Category == category {
					Todelete = append(Todelete, n)
					isCtg = true
				}
			}
		}
		return nil, Todelete, isCtg
	}
	return nil, nil, false
}
