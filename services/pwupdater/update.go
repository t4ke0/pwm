package pwupdater

import (
	"github.com/TaKeO90/pwm/services/pwshow"
	"github.com/TaKeO90/pwm/services/serverenc"
	"github.com/TaKeO90/pwm/sqlite"
)

//CredsUpdate struct holds pwid of the credentials and credentials
type CredsUpdate struct {
	ID          int
	Credentials pwshow.UserStuff
}

//CredUpdateList list of CredsUpdate
type CredUpdateList []CredsUpdate

//ParseMap parse map that we get from comparing front-end data and back-end's
func ParseMap(m map[pwshow.UserStuff]pwshow.UserStuff, filter bool) CredUpdateList {
	var res []CredsUpdate
	c := &CredsUpdate{}

	for k, v := range m {
		pwid := pwshow.GetPWID(k.Username, k.Password, k.Category)
		if pwid != 0 {
			c.ID = pwid
			if k.Username != v.Username {
				c.Credentials.Username = v.Username
			}
			if k.Password != v.Password {
				c.Credentials.Password = v.Password
			}
			if k.Category != v.Category && !filter {
				c.Credentials.Category = v.Category
			}
			res = append(res, *c)
		}
	}
	return res
}

//UpdateCreds update user credentials
func (cl CredUpdateList) UpdateCreds(username string) (bool, error) {
	db := sqlite.InitDb()
	var ok bool
	for _, n := range cl {
		Hpw, err := serverenc.CredReveal(username, n.Credentials.Password, "", true)
		if err != nil {
			return false, err
		}
		uOk, pOk, cOk := sqlite.Update(n.ID, db, n.Credentials.Username, Hpw, n.Credentials.Category)
		if uOk || pOk || cOk {
			ok = true
		} else {
			ok = false
			break
		}
	}
	return ok, nil
}
