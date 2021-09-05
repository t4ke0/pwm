package main

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go/v4"
)

type tokenClaims struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	SessionID    string `json:"session_id"`
	SymmetricKey string `json:"symmetrickey"`
}

func getNewJWTtoken(authServerKey []byte, claims tokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":       claims.UserID,
		"username":     claims.Username,
		"sessionid":    claims.SessionID,
		"symmetrickey": claims.SymmetricKey,
	})
	return token.SignedString(authServerKey)
}

func parseJWTtoken(tokenString string, authServerKey []byte) (tokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return authServerKey, nil
	})

	if err != nil {
		return tokenClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return tokenClaims{}, fmt.Errorf("type assertion error token claims")
	}
	if !token.Valid {
		return tokenClaims{}, fmt.Errorf("token not valid")
	}
	log.Printf("debug: %v", claims)
	return tokenClaims{
		UserID:       int(claims["userid"].(float64)),
		Username:     claims["username"].(string),
		SessionID:    claims["sessionid"].(string),
		SymmetricKey: claims["symmetrickey"].(string),
	}, nil
}
