package main

import (
	"encoding/hex"
	"fmt"

	"./pwencrypter"
)

//TODO Generate A key for the Server to encrypt all users keys and decrypt them with it
// Save that key in a safe place

type User struct {
	username string
	password string
}

type Encrypted struct {
	password_Encrypted []byte
	password_Decrypted string
}

var (
	u User
	e Encrypted
)

func main() {
	u.username = "yassine"
	u.password = "Secure"

	genkey := pwencrypter.GenKeyP(u.password)
	if ok := pwencrypter.SaveKey(genkey, u.username); ok {
		fmt.Println("Saved SuccessFully")
	}
	key := pwencrypter.LoadKey(u.username)

	enc := pwencrypter.Encrypt(u.password, key)
	e.password_Encrypted = enc
	hexenc := make([]byte, hex.EncodedLen(len(e.password_Encrypted)))
	hex.Encode(hexenc, enc)
	HEX := (string(hexenc))
	fmt.Println([]byte(HEX))
	fmt.Printf("Encrypted :%s\n", string(hexenc))

	dec := pwencrypter.Decrypt(enc, key)
	e.password_Decrypted = dec
	fmt.Printf("Decrypted : %s\n", e.password_Decrypted)
}
