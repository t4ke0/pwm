package serverenc

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"log"

	"../pwencrypter"
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
	//TODO Look For A place Where to hide this key
	return isSaved
}

// LookForServerKey Search for server encrpytion key if found return true otherwise return false
func LookForServerKey() bool {
	var found bool
	files, err := ioutil.ReadDir(KeysDir)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) != 0 {
		for _, f := range files {
			if f.Name() == "server.key" {
				found = true
				break
			} else {
				found = false
			}
		}
	}
	return found
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
