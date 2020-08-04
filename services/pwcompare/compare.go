package pwcompare

import (
	"log"

	"github.com/TaKeO90/pwm/services/pwdelete"
	"github.com/TaKeO90/pwm/services/pwsaver"
	"github.com/TaKeO90/pwm/services/pwshow"
	"github.com/TaKeO90/pwm/services/pwupdater"
)

// Compare function that compares two json arrays one of them
// is back-end data and the other one is the front-end's
func Compare(a1, a2 pwshow.UserList) (index []int) {
	m := make(map[pwshow.UserStuff]bool)
	var (
		arr1 pwshow.UserList
		arr2 pwshow.UserList
	)
	if len(a1) > len(a2) {
		arr2, arr1 = a1, a2
	} else {
		arr1, arr2 = a1, a2
	}
	for _, item := range arr1 {
		m[item] = true
	}
	for i, item := range arr2 {
		if _, ok := m[item]; !ok {
			index = append(index, i)
		}
	}
	return
}

// DiffAnalyse function that analyse the diffrence we got from the function above to check if we need to
// update , save or delete credentials
func DiffAnalyse(username, category string, a1, a2 pwshow.UserList, index []int) (ok bool) {
	var orDiff bool
	var filter bool

	if category != "" {
		filter = true
	} else {
		filter = false
	}

	for _, i := range index {
		if len(a1) > i && len(a2) > i {
			orDiff = (a1[i].Username != a2[i].Username || a1[i].Password != a2[i].Password || a1[i].Category != a2[i].Category)
		}
		if len(a1) < i+1 {
			//			fmt.Println("saving", a2[i])
			ok = pwsaver.AddCreds(a2[i].Username, a2[i].Password, a2[i].Category, username)

		} else if len(a1) > i && orDiff {
			updateM := make(map[pwshow.UserStuff]pwshow.UserStuff)
			updateM[a1[i]] = a2[i]
			check := checkForCreds(a2[i], a1)
			if !check {
				cl := pwupdater.ParseMap(updateM, filter)
				isUp, err := cl.UpdateCreds(username)
				ok = isUp
				if err != nil {
					log.Fatal(err)
				}
				//				fmt.Println("update", updateM)
			} else {
				isDel, err := pwdelete.DeleteCreds(a1[i].Username, a1[i].Password, a1[i].Category, filter)
				if err != nil {
					log.Fatal(err)
				}
				ok = isDel
			}
		} else {
			//			fmt.Println("delete", a1[i])
			isDel, err := pwdelete.DeleteCreds(a1[i].Username, a1[i].Password, a1[i].Category, filter)
			if err != nil {
				log.Fatal(err)
			}
			ok = isDel
		}
	}
	return
}

func checkForCreds(cred pwshow.UserStuff, a1 pwshow.UserList) (ok bool) {
	for _, item := range a1 {
		if cred.Username == item.Username && cred.Password == item.Password && cred.Category == item.Category {
			ok = true
		} else {
			ok = false
		}
	}
	return
}
