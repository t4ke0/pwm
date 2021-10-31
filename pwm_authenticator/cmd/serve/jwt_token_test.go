package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"testing"
)

var testClaims tokenClaims = tokenClaims{
	UserID:       69,
	Username:     "pwm_user",
	SessionID:    "pwm_test_session_id",
	SymmetricKey: "pwm_test_user_key",
}

var pwmTestServerKey []byte = []byte("pwm_server_key")

var testJWTtoken string

func getHashOfClaims(c tokenClaims) (string, error) {
	buff := new(bytes.Buffer)
	if err := gob.NewEncoder(buff).Encode(c); err != nil {
		return "", err
	}

	return hex.EncodeToString(sha256.New().Sum(buff.Bytes())), nil
}

func TestGetNewJWTtoken(t *testing.T) {
	var err error
	testJWTtoken, err = getNewJWTtoken(pwmTestServerKey, testClaims)
	if err != nil {
		t.Logf("failed to get JWT token (%v)", err)
		t.Fail()
		return
	}
}

func TestGetClaimsFromJWTtoken(t *testing.T) {
	claims, err := parseJWTtoken(testJWTtoken, pwmTestServerKey)
	if err != nil {
		t.Logf("failed to parse JWT token (%v)", err)
		t.Fail()
		return
	}

	h1, err := getHashOfClaims(claims)
	if err != nil {
		t.Logf("%v", err)
		t.Fail()
		return
	}

	h2, err := getHashOfClaims(testClaims)
	if err != nil {
		t.Logf("%v", err)
		t.Fail()
		return
	}

	if h1 != h2 {
		t.Logf("failed to parse JWT token [%v != %v]", h1, h2)
		t.Fail()
		return
	}
}
