package pwencrypter

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/scrypt"
)

const (
	//SALTBYTE salt size
	SALTBYTE = 32
	//KeysPath server and users encryption key path
	KeysPath = "./services/pwencrypter/keys"
)

// CheckError Checks if err is not nil then it logs the error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GenKeyP func generate an encryption key from user password
// Takes password as an input and returns byte slice
func GenKeyP(p string) []byte {
	salt := make([]byte, SALTBYTE)
	key, err := scrypt.Key([]byte(p), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

// SaveKey Func Save user's encryption key into a file named with his username
// The Func takes k encryption key the user generate in the beginning and user "username of the user"
func SaveKey(k []byte, user string) bool {
	filename := filepath.Join(KeysPath, user+".key")
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	CheckError(err)
	defer f.Close()
	if _, err := f.Write(k); err != nil {
		log.Fatal(err)
	}
	return true
}

// LoadKey Func Reads the key of each user from the keys folder to use it in encryption or decryption
// This Function needs Username as an input & it returns slice of byte which is the key
func LoadKey(user string) []byte {
	b := make([]byte, 64)
	filename := filepath.Join(KeysPath, user+".key")
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	CheckError(err)
	defer f.Close()
	n, err := f.Read(b)
	CheckError(err)
	return b[:n]
}

// Encrypt Func encrypt passwords of the user basing on a key that the user have generated before
// This Function takes password as a string type & key = k as a []byte type as inputs
// returns encrypted password as byte array
func Encrypt(password string, k []byte) []byte {
	c, err := aes.NewCipher(k)
	CheckError(err)
	gcm, err := cipher.NewGCM(c)
	CheckError(err)
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		log.Fatal(err)
	}
	pw := []byte(password)
	f, _ := gcm.Seal(nonce, nonce, pw, nil), 0
	return f
}

// Decrypt Func decrypt passwords of the user basing on the user key
// This Function takes ciphertext which is the encrypted password & key = k as inputs
// returns the decrypted password as a string
func Decrypt(ciphertext, k []byte) string {
	c, err := aes.NewCipher(k)
	CheckError(err)
	gcm, err := cipher.NewGCM(c)
	CheckError(err)
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Fatal(errors.New("ciphertext too short"))
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	f, err := gcm.Open(nil, nonce, ciphertext, nil)
	CheckError(err)
	return string(f)
}
