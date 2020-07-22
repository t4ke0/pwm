package serverenc

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"
	"regexp"

	"github.com/TaKeO90/pwm/services/pwencrypter"
	"github.com/TaKeO90/pwm/sqlite"
)

// KeysDir server key directory
const KeysDir string = "./services/pwencrypter/keys"

// GenerateRandomPassword for The Server for Generating an encryption key
func GenerateRandomPassword() string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	// Random password length is 15
	b := make([]byte, 15)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	for i, j := range b {
		b[i] = letters[j%byte(len(letters))]
	}
	return string(b)
}

// GenerateServerKey generate server's encryption key
// save the Generated Key into server encrytion key folder
func GenerateServerKey() bool {
	pwd := GenerateRandomPassword()
	key := pwencrypter.GenKeyP(pwd)
	isSaved := pwencrypter.SaveKey(key, "server")
	//TODO figure out how to hide server's encryption key
	return isSaved
}

// SearchForKeys search for server's key and also user's keys
func SearchForKeys(server, users bool) (bool, error) {
	var amountOfuser int
	files, err := ioutil.ReadDir(KeysDir)
	if err != nil {
		return false, err
	}
	if len(files) != 0 {
		for _, f := range files {
			if server {
				if f.Name() == "server.key" {
					return true, nil
				}
			} else if users {
				v := regexp.MustCompile(`\w+.key`)
				if matched := v.MatchString(f.Name()); matched && f.Name() != "server.key" {
					amountOfuser++
				}
			}
		}
		db := sqlite.InitDb()
		amount, err := sqlite.CountUsers(db)
		if err != nil {
			return false, err
		}
		if amountOfuser == amount {
			return true, nil
		}
	}
	return false, nil
}

// CheckError check for errors then log error if err is not nil
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// EncryptUserKey encrypt users key
func EncryptUserKey(userkey, key []byte) []byte {
	encK := make([]byte, hex.EncodedLen(len(userkey)))
	hex.Encode(encK, userkey)
	encrk := pwencrypter.Encrypt(string(encK), key)
	return encrk
}

// DecryptUserKey decrypt users key
func DecryptUserKey(userkey, key []byte) []byte {
	decK := pwencrypter.Decrypt(userkey, key)
	hexdec := make([]byte, hex.DecodedLen(len([]byte(decK))))
	n, err := hex.Decode(hexdec, []byte(decK))
	CheckError(err)
	return hexdec[:n]
}

//CredReveal Encrypt or Decrypt user's passwords
func CredReveal(username, password, decPW string, enc bool) (string, error) {
	key, err := ioutil.ReadFile(path.Join(KeysDir, username+".key"))
	if err != nil {
		return "", err
	}
	//Load server key
	serverK := pwencrypter.LoadKey("server")
	decKey := DecryptUserKey(key, serverK)
	if enc && password != "" {
		// Decrypt user key for encrypting his password
		Epw := pwencrypter.Encrypt(password, decKey)
		hexenc := make([]byte, hex.EncodedLen(len(Epw)))
		hex.Encode(hexenc, Epw)
		pwH := string(hexenc)
		return pwH, nil
	} else if !enc && decPW != "" {
		src := []byte(decPW)
		dst := make([]byte, hex.DecodedLen(len(src)))
		n, err := hex.Decode(dst, src)
		if err != nil {
			return "", err
		}
		pwH := pwencrypter.Decrypt(dst[:n], decKey)
		return pwH, nil
	}
	return "", nil
}
