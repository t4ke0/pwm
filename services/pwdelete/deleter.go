package pwdelete

import (
	"github.com/TaKeO90/pwm/services/pwshow"
	"github.com/TaKeO90/pwm/sqlite"
)

//DeleteCreds delete credentials
func DeleteCreds(creduser, password, category string, isctg bool) (bool, error) {
	db := sqlite.InitDb()
	pwid := pwshow.GetPWID(creduser, password, category)
	if isctg {
		if ok, err := sqlite.Delete(pwid, category, db); ok {
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
