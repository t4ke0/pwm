package genkey

import (
	"fmt"

	"github.com/TaKeO90/pwm/services/serverenc"
)

func KeysChecking() (string, error) {

	if found, err := serverenc.SearchForKeys(true, false); found {
		if err != nil {
			return "", err
		}
		return "Server's Key found", nil
	} else {
		// search for users encryption key if there is no user keys then we through an err
		// otherwise we need to create a server key
		usersKfound, err := serverenc.SearchForKeys(false, true)
		if err != nil {
			return "", err
		}
		if !usersKfound {
			//we generate new server key
			if isOk := serverenc.GenerateServerKey(); isOk {
				return "Generated a Server Key", nil
			}
		} else {
			return "", fmt.Errorf("Server's key not found\n")
		}
	}
	return "", nil
}
