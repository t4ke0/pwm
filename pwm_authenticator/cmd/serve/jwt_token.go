package main

import "github.com/dgrijalva/jwt-go"

// TokenClaims is the information that a JWT token will hold.
type TokenClaims struct {
	Username     string
	SessionID    string
	SymmetricKey string
}

// GetNewJWTtoken get a signed JWT token.
func GetNewJWTtoken(authServerKey []byte, claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":     claims.Username,
		"sessionID":    claims.SessionID,
		"SymmetricKey": claims.SymmetricKey,
	})
	return token.SignedString(authServerKey)
}
