package pwdelete

import (
	"../../sqlite"
)

// This Func Deletes unwanted creds
func DeleteCreds(id int) bool {
	var isOk bool
	db := sqlite.InitDb()
	if ok := sqlite.Delete(id, db); ok {
		isOk = true
	} else {
		isOk = false
	}
	return isOk
}
