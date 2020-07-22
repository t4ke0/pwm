package pwdelete

import (
	"github.com/TaKeO90/pwm/services/pwshow"
	"github.com/TaKeO90/pwm/sqlite"
)

//DeleteCreds delete credentials
func DeleteCreds(delList pwshow.UserList, isctg bool) (bool, error) {
	db := sqlite.InitDb()
	for _, n := range delList {
		pwid := pwshow.GetPWID(n.Username, n.Password, n.Category)
		if isctg {
			if ok, err := sqlite.Delete(pwid, n.Category, db); ok {
				if err != nil {
					return false, err
				}
				return true, nil
			}
			return false, nil
		}
		if ok, err := sqlite.Delete(pwid, "", db); ok {
			if err != nil {
				return false, nil
			}
			return true, nil
		}
		return false, nil
	}
	return false, nil
}
