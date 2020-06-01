package main

import (
	"../../services/serverenc"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// KeyGen Generate the server encryption key
func KeyGen() {
	if isGen := serverenc.GenerateServerKey(); isGen {
		fmt.Println("Generated Server encryption key")
	}
}

// CopyFile We use this to copy server encryption key if it's not in the Right place When launching the launcher
func CopyFile(src, dst string) bool {
	readSrc, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	//Write src into dst
	if err := ioutil.WriteFile(dst, readSrc, 0600); err != nil {
		log.Fatal(err)
	}
	return true
}

func main() {
	const dstP string = "./services/pwencrypter/keys/server.key"

	if found := serverenc.LookForServerKey(); found {
		fmt.Println("Server's Key found")
	} else {
		fmt.Printf("Have You Changed the Path of the server key ? [[Y]/n] ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		newKeyPath := bufio.NewScanner(os.Stdin)
		switch scanner.Text() {
		case "n":
			KeyGen()
		case "y":
			fmt.Printf("Enter key path : ")
			newKeyPath.Scan()
			if isCopied := CopyFile(newKeyPath.Text(), dstP); isCopied {
				fmt.Println("Copied The Key SuccessFuly")
			}
		default:
			fmt.Printf("Enter key path : ")
			newKeyPath.Scan()
			if isCopied := CopyFile(newKeyPath.Text(), dstP); isCopied {
				fmt.Println("Copied The Key SuccessFuly")
			}
		}
	}
}
