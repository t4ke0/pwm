package pwdelete

import (
	"../../sqlite"
)

//DeleteCreds This Func Deletes unwanted creds
func DeleteCreds(username string) bool {
	var isOk bool
	db := sqlite.InitDb()
	uid := sqlite.GetUID(username, db)
	if ok := sqlite.Delete(uid, db); ok {
		isOk = true
	} else {
		isOk = false
	}
	return isOk
}
