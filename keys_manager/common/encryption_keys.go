package common

import (
	"crypto/aes"
	"crypto/cipher"
	random "crypto/rand"
	"encoding/hex"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const KeyLength int = 32

func readWordFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return []string{}, err
	}
	content := strings.Split(string(data), "\n")
	return content, nil
}

func gcmEncryption(key, src []byte) ([]byte, error) {
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(random.Reader, nonce); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := aesgcm.Seal(nonce, nonce, src, nil)
	return ciphertext, nil
}

func gcmDecrpytion(key, encryptedtext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := encryptedtext[:nonceSize], encryptedtext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

type Key []byte

func GenerateEncryptionKey(filePath string, size int) (Key, error) {
	if size == 0 {
		size = KeyLength
	}
	fileContent, err := readWordFile(filePath)
	if err != nil {
		return nil, err
	}
	key := make(Key, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		word := fileContent[rand.Intn(len(fileContent))]
		var l byte
		if len(word) >= 2 {
			l = word[rand.Intn(2)]
		}
		key[i] = l
	}
	return key, nil
}

func (k Key) String() string {
	return hex.EncodeToString(k)
}

// Encrypt encrypts the user key using servers key.
func (k Key) Encrypt(plainUserKey Key) ([]byte, error) {
	return gcmEncryption(k, plainUserKey)
}

// Decrypt decrypts the user key using servers key.
func (k Key) Decrypt(encryptedUserKey []byte) (Key, error) {
	decrypted, err := gcmDecrpytion(k, encryptedUserKey)
	if err != nil {
		return nil, err
	}
	return Key(decrypted), nil
}
