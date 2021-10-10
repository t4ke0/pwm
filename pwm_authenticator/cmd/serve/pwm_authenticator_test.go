package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/t4ke0/pwm/pwm_authenticator/api"
)

type httpMethod = string

const (
	GET  httpMethod = "GET"
	POST            = "POST"
)

// NOTE: at the moment this unit testing is going to be running only on
// localhost not on CI.

func makeTestHTTPrequest(method httpMethod, url string, data []byte, headerValue map[string]string) (*httptest.ResponseRecorder, error) {

	setupGinEngine()

	var buff *bytes.Buffer
	if data == nil {
		buff = &bytes.Buffer{}
	} else {
		buff = bytes.NewBuffer(data)
	}

	r, err := http.NewRequest(method, url, buff)
	if err != nil {
		return nil, err
	}

	if headerValue != nil {
		for k, v := range headerValue {
			r.Header.Set(k, v)
		}
	}

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, r)

	return w, nil
}

const (
	testUsername api.Field = "yassine.test"
	testPassword api.Field = "testpassword"
	testEmail    api.Field = "test@test.com"
)

var tokenFromAuth string

func TestRegistration(t *testing.T) {

	requestBody := api.RegisterRequest{
		Username: testUsername,
		Password: testPassword,
		Email:    testEmail,
	}
	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Logf("Error: marshal json body %v", err)
		t.Fail()
		return
	}
	response, err := makeTestHTTPrequest(POST, "/register", data, nil)
	if err != nil {
		t.Logf("Error body: %v", response.Body.String())
		t.Logf("Error: register POST request %v", err)
		t.Fail()
		return
	}

	if response.Code != http.StatusCreated {
		t.Logf("Error: response not equal to 201 (%v) | (%v)", response.Code, response.Body.String())
		t.Fail()
		return
	}

}

func TestLogin(t *testing.T) {

	requestBody := api.AuthRequest{
		Username: testUsername,
		Password: testPassword,
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		t.Logf("Error marshal JSON data  %v", err)
		t.Fail()
		return
	}

	response, err := makeTestHTTPrequest(POST, "/login", data, nil)
	if err != nil {
		t.Logf("Error: request /login %v", err)
		t.Fail()
		return
	}

	if response.Code != http.StatusOK {
		t.Logf("Failed to login %v", response.Code)
		t.Fail()
		return
	}

	var token api.AuthResponse
	if err := json.Unmarshal(response.Body.Bytes(), &token); err != nil {
		t.Logf("Failed tp unmarshal JSON response data %v", err)
		t.Fail()
		return
	}

	tokenFromAuth = token.JwtToken

	t.Logf("JWT token %v", tokenFromAuth)
}

func TestGetJWTinfo(t *testing.T) {

	response, err := makeTestHTTPrequest(GET, "/info", nil, map[string]string{
		"token": tokenFromAuth,
	})

	if err != nil {
		t.Logf("failed to make info request %v", err)
		t.Fail()
		return
	}

	if response.Code != http.StatusOK {
		t.Logf("expected 200 got(%v) [%v]", response.Code, response.Body.String())
		t.Fail()
		return
	}

	t.Logf("token claims %v", response.Body.String())
}
