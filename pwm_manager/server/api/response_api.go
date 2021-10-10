package api

// TokenClaims JSON data returned by `/info` endpoint of the
// `pwm_authenticator`.
type TokenClaims struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	SessionID    string `json:"session_id"`
	SymmetricKey string `json:"symmetrickey"`
}
