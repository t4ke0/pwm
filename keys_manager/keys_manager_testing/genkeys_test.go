package keys_manager_testing

import (
	"testing"

	"github.com/t4ke0/pwm/keys_manager/common"
)

const wordsFilePath = "../common/words.txt"

var key, userKey common.Key

var encryptedKey []byte

func TestGenerateKey(t *testing.T) {
	var err error
	key, err = common.GenerateEncryptionKey(wordsFilePath, 0)
	if err != nil {
		t.Logf("Generate server Key [%v]", err)
		t.Fail()
		return
	}
	t.Logf("Generated server key [%v]", key.String())

	userKey, err = common.GenerateEncryptionKey(wordsFilePath, 0)
	if err != nil {
		t.Logf("Generate user key [%v]", err)
		t.Fail()
		return
	}
	t.Logf("Generated User key [%v]", userKey.String())
}

func TestEncryptUserKey(t *testing.T) {
	if key == nil {
		t.Log("key is nil")
		t.Fail()
		return
	}
	var err error
	encryptedKey, err = key.Encrypt(userKey)
	if err != nil {
		t.Logf("encrypt user key [%v]", err)
		t.Fail()
		return
	}
	t.Logf("encrypted User key [%x]", encryptedKey)
}

func TestDecrpytUserKey(t *testing.T) {
	if encryptedKey == nil {
		t.Log("key is nil")
		t.Fail()
		return
	}
	decryptedKey, err := key.Decrypt(encryptedKey)
	if err != nil {
		t.Logf("decrypt user key [%v]", err)
		t.Fail()
		return
	}
	if decryptedKey.String() != userKey.String() {
		t.Log("user key is not the same after decypting.")
		t.Fail()
		return
	}
	t.Logf("decrypted User key [%v]", decryptedKey.String())
}
