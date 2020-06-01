package pwdelete

//TODO: support deleting by category
import (
	"../../sqlite"
)

//DeleteCreds This Func Deletes unwanted creds
func DeleteCreds(username, category string) bool {
	var isOk bool
	db := sqlite.InitDb()
	if username != "" && category == "" {
		uid := sqlite.GetUID(username, db)
		if ok := sqlite.Delete(uid, category, db); ok {
			isOk = true
		} else {
			isOk = false
		}
	} else if category != "" && username != "" {
		uid := sqlite.GetUID(username, db)
		if ok := sqlite.Delete(uid, category, db); ok {
			isOk = true
		} else {
			isOk = false
		}
	}
	return isOk
}
