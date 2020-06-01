package pwdelete

import (
	"../../sqlite"
)

// This Func Deletes unwanted creds
// TODO: userid instead of the id of the credentials here
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
